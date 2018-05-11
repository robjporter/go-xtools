package xipfilter

import (
	"net"
	"sync"
)

type Options struct {
	//explicity allowed IPs
	AllowedIPs []string
	//explicity blocked IPs
	BlockedIPs []string
	//block by default (defaults to allow)
	BlockByDefault bool
}

// IPFilter ...
type IPFilter struct {
	opts Options
	//mut protects the below
	//rw since writes are rare
	mut            sync.RWMutex
	defaultAllowed bool
	ips            map[string]bool
	subnets        []*subnet
}

type subnet struct {
	str     string
	ipnet   *net.IPNet
	allowed bool
}

//New ...
func New(opts Options) *IPFilter {
	f := &IPFilter{
		opts:           opts,
		ips:            map[string]bool{},
		defaultAllowed: !opts.BlockByDefault,
	}

	for _, ip := range opts.BlockedIPs {
		f.BlockIP(ip)
	}

	for _, ip := range opts.AllowedIPs {
		f.AllowIP(ip)
	}
	return f
}

// AllowIP ...
func (f *IPFilter) AllowIP(ip string) bool {
	return f.ToggleIP(ip, true)
}

// BlockIP ..
func (f *IPFilter) BlockIP(ip string) bool {
	return f.ToggleIP(ip, false)
}

// ToggleIP ...
func (f *IPFilter) ToggleIP(str string, allowed bool) bool {
	//check if has subnet
	if ip, ipNet, err := net.ParseCIDR(str); err == nil {
		// containing only one ip?
		if n, total := ipNet.Mask.Size(); n == total {
			f.mut.Lock()
			f.ips[ip.String()] = allowed
			f.mut.Unlock()
			return true
		}
		//check for existing
		f.mut.Lock()
		found := false
		for _, subnet := range f.subnets {
			if subnet.str == str {
				found = true
				subnet.allowed = allowed
				break
			}
		}
		if !found {
			f.subnets = append(f.subnets, &subnet{
				str:     str,
				ipnet:   ipNet,
				allowed: allowed,
			})
		}
		f.mut.Unlock()
		return true
	}
	//check if plain ip
	if ip := net.ParseIP(str); ip != nil {
		f.mut.Lock()
		f.ips[ip.String()] = allowed
		f.mut.Unlock()
		return true
	}
	return false
}

//Allowed returns if a given IP can pass through the filter
func (f *IPFilter) Allowed(ipstr string) bool {
	return f.NetAllowed(net.ParseIP(ipstr))
}

//NetAllowed returns if a given net.IP can pass through the filter
func (f *IPFilter) NetAllowed(ip net.IP) bool {
	//invalid ip
	if ip == nil {
		return false
	}
	//read lock entire function
	//except for db access
	f.mut.RLock()
	defer f.mut.RUnlock()
	//check single ips
	allowed, ok := f.ips[ip.String()]
	if ok {
		return allowed
	}
	//scan subnets for any allow/block
	blocked := false
	for _, subnet := range f.subnets {
		if subnet.ipnet.Contains(ip) {
			if subnet.allowed {
				return true
			}
			blocked = true
		}
	}
	if blocked {
		return false
	}
	//use default setting
	return f.defaultAllowed
}

//Blocked returns if a given IP can NOT pass through the filter
func (f *IPFilter) Blocked(ip string) bool {
	return !f.Allowed(ip)
}

//NetBlocked returns if a given net.IP can NOT pass through the filter
func (f *IPFilter) NetBlocked(ip net.IP) bool {
	return !f.NetAllowed(ip)
}

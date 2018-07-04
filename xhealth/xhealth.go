package xhealth

import (
	"net"
)

func New() *Pinger {
	return &Pinger{
		Running:  false,
		Monitors: make(map[string]*monitor),
	}
}

func (p *Pinger) isMonitor(name string) bool {
	for k, _ := range p.Monitors {
		if k == name {
			return true
		}
	}
	return false
}

func (p *Pinger) StartMonitors() {
	if !p.Running {
		for k, _ := range p.Monitors {
			go p.Monitors[k].run()
			p.Monitors[k].active = true
		}
		p.Running = true
	}
}

func (p *Pinger) StopMonitors() {
	if p.Running {
		for k, _ := range p.Monitors {
			p.Monitors[k].active = false
		}
		p.Running = false
	}
}

func (p *Pinger) ResolveHost(name string) (net.IP, error) {
	return getIPAddr(name)
}

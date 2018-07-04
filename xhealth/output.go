package xhealth

import (
	"time"
)

func (p *Pinger) Count() int {
	return len(p.Monitors)
}

func (p *Pinger) ActiveCount() int {
	count := 0
	for k, _ := range p.Monitors {
		if p.Monitors[k].active {
			count++
		}
	}
	return count
}

func (p *Pinger) GetMonitorNames() []string {
	var names []string
	for k, _ := range p.Monitors {
		names = append(names, k)
	}
	return names
}

func (p *Pinger) LastCheck(name string) int64 {
	return p.Monitors[name].lastcheck
}

func (p *Pinger) LastStatus(name string) bool {
	return p.Monitors[name].laststatus
}

func (p *Pinger) LastDuration(name string) time.Duration {
	return p.Monitors[name].lastduration
}

func (p *Pinger) AverageDuration(name string) time.Duration {
	return time.Duration(float64(p.Monitors[name].totalduration) / float64(p.Monitors[name].totalchecks))
}

func (p *Pinger) TotalSuccess(name string) uint64 {
	return p.Monitors[name].totalsuccess
}

func (p *Pinger) TotalFails(name string) uint64 {
	return p.Monitors[name].totalfails
}

func (p *Pinger) TotalChecks(name string) uint64 {
	return p.Monitors[name].totalchecks
}

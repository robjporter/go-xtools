package xhealth

import (
	"time"
)

func (m *monitor) run() {
	go func() {
		for {
			if m.active {
				m.lastcheck = time.Now().Unix()
				time.Sleep(time.Duration(m.frequency) * time.Second)
				var err error
				var dur time.Duration
				var resp string

				switch m.monitorType {
				case URL:
					dur, resp, err = m.callHTTP()
				case PING:
					dur, resp, err = m.callPing()
				}

				previousState := m.laststatus
				m.totalduration += dur
				m.lastduration = dur
				m.lastresponse = resp
				if err != nil {
					m.laststatus = false
					m.totalfails++
				} else {
					m.laststatus = true
					m.totalsuccess++
				}

				if m.laststatus != previousState && m.changedcallback != nil {
					m.changedcallback(m.name, m.lastcheck, m.laststatus, m.lastresponse)
				}

				m.totalchecks++

				if m.callback != nil {
					m.callback(m.name, m.lastcheck, m.laststatus, m.lastresponse)
				}
			}
		}
	}()
}

func (p *Pinger) Callback(name string, callback func(name string, check int64, status bool, response string)) {
	if p.isMonitor(name) {
		p.Monitors[name].callback = callback
	}
}

func (p *Pinger) ChangedCallback(name string, callback func(name string, check int64, status bool, response string)) {
	if p.isMonitor(name) {
		p.Monitors[name].changedcallback = callback
	}
}

package xhealth

import (
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)

func (p *Pinger) NewURLMonitor(name, url string, frequency int) {
	if !p.isMonitor(name) {
		p.Monitors[name] = &monitor{
			active:      false,
			name:        name,
			url:         url,
			monitorType: URL,
			frequency:   frequency,
			laststatus:  false,
			totalchecks: 0,
		}
	}
}

func (m monitor) callHTTP() (time.Duration, string, error) {
	start := time.Now()
	response, err := http.Get(m.url)
	defer response.Body.Close()
	if err != nil {
		return 0, "", err // errors.WithStack(err)
	}
	if response.StatusCode >= 400 {
		return 0, "", fmt.Errorf("Invalid state %s", response.Status)
	}
	bodyBytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return 0, "", err2
	}
	bodyString := string(bodyBytes)
	return time.Now().Sub(start), bodyString, nil
}

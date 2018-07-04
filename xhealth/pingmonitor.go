package xhealth

import (
	"time"
	"os/exec"
	"bytes"
	"regexp"
	"errors"
	"strings"
	"strconv"
)

func (p *Pinger) NewPingMonitor(name, url string, frequency int) {
	if !p.isMonitor(name) {
		p.Monitors[name] = &monitor{
			active:      false,
			name:        name,
			url:         url,
			frequency:   frequency,
			laststatus:  false,
			monitorType: PING,
			totalchecks: 0,
		}
	}
}

func (m monitor) callPing() (time.Duration, string, error) {
	cmd := exec.Command("ping", m.url, "-c 3")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	if errStr == "" && err == nil {
		loss := regexp.MustCompile(`100.0% packet loss`)
		times := regexp.MustCompile(`time=(.*)`)
		if loss.MatchString(outStr) {
			return 0, "", errors.New("Unable to connect")
		}
		if times.MatchString(outStr) {
			matches := times.FindAllStringSubmatch(outStr, -1)
			total := 0.0
			if len(matches) > 0 {
				for i := 0; i < len(matches); i++ {
					if len(matches[i]) > 1 {
						total += getPingTime(matches[i][1])
					}
				}
			}
			return time.Duration(time.Duration(total/3) * time.Millisecond), outStr, nil
		}
	}
	return 0, "", errors.New("Unable to connect")
}

func getPingTime(time string) float64 {
	splits := strings.Split(time, " ")
	if len(splits) > 1 {
		a, err := strconv.ParseFloat(splits[0], 64)
		if err == nil {
			return a
		}
	}
	return 0
}

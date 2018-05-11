package xtimer

import (
	"strings"
	"time"
)

var timers map[string]time.Time

func init() {
	timers = make(map[string]time.Time)
}

func Timer(name string) string {
	name = strings.TrimSpace(name)

	if _, ok := timers[name]; ok {
		end := time.Now().Round(time.Second)
		return end.Sub(timers[name]).String()
	} else {
		timers[name] = time.Now().Round(time.Second)
	}

	return ""
}

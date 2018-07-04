package xhealth

import (
	"time"
)

type Type uint8

const (
	URL     Type = 1 + iota
	PING
	CURL
	POST
	CONTENT
)

type monitor struct {
	active          bool
	name            string
	url             string
	frequency       int
	monitorType     Type
	totalchecks     uint64
	totalsuccess    uint64
	totalfails      uint64
	lastcheck       int64
	laststatus      bool
	lastresponse    string
	lastduration    time.Duration
	totalduration   time.Duration
	callback        func(name string, check int64, status bool, response string)
	changedcallback func(name string, check int64, status bool, response string)
}

type Pinger struct {
	Running  bool
	Monitors map[string]*monitor
}

package xhttp

// TODO

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"sync"
)

// url is the URL to hit
var url string

// done is used to report completion to main thread
var done chan bool
var mutex = &sync.Mutex{}

type Hammer struct {
	threadCount uint
	hitCount uint
	delay time.Duration
	spread time.Duration
	url string
	start time.Time
	quickest time.Duration
	slowest time.Duration
	failed int
	hammered uint
	duration time.Duration
}

func NewHammer() *Hammer {
	return &Hammer{
		threadCount:uint(4),
		hitCount:uint(20),
		delay:100*time.Millisecond,
		spread:1*time.Second,
		quickest: time.Duration(10*time.Second),
		slowest: 0,
		hammered:0,
	}
}

func (h *Hammer) SetThreadCount(number uint) *Hammer {
	h.threadCount=number
	return h
}

func (h *Hammer) SetHitCount(number uint) *Hammer {
	h.hitCount=number
	return h
}

func (h *Hammer) SetDelay(delay time.Duration) *Hammer {
	h.delay=delay
	return h
}

func (h *Hammer) SetSpread(spread time.Duration) *Hammer {
	h.spread=spread
	return h
}

func (h *Hammer) SetURL(url string) *Hammer {
	h.url = url
	return h
}

func (h *Hammer) GetStartTime() time.Time {
	return h.start
}

func (h *Hammer) GetDuration() time.Duration {
	return h.duration
}

func (h *Hammer) GetQuickest() time.Duration {
	return h.quickest
}

func (h *Hammer) GetSlowest() time.Duration {
	return h.slowest
}

func (h *Hammer) GetFailed() int {
	return h.failed
}

func (h *Hammer) GetHits() uint {
	return h.hammered
}

// hitURL hits a URL and reads the result
func (h *Hammer) hitURL() {
	sta := time.Now()
	resp, err := http.Get(h.url)
	sto := time.Now()
	fail := 0
	if err != nil {
		log.Printf("Error getting URL: %v", err)
		fail += 1
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		fail += 1
	}
	h.process(sta,sto,fail)
}

func (h *Hammer) process(start,stop time.Time,fail int) {
	mutex.Lock()
	h.hammered++
	h.failed += fail
	taken := stop.Sub(start)
	if taken < h.quickest {h.quickest=taken}
	if taken > h.slowest {h.slowest=taken}
	mutex.Unlock()
}

// hammer repeatedly hits a URL (used as a goroutine)
func (h *Hammer) hammer() {
	unlimited := h.hitCount == 0

	for i := uint(0); i < h.hitCount || unlimited; i++ {
		h.hitURL()
		time.Sleep(h.delay)
	}

	done <- true
}

// main
func (h *Hammer) Run() {
	done = make(chan bool)

	// Launch all threads
	h.start = time.Now()
	for i := uint(0); i < h.threadCount; i++ {
		go h.hammer()

		if h.spread > 0 {
			time.Sleep(h.spread / time.Duration(h.threadCount))
		}
	}
	h.duration = time.Now().Sub(h.start)

	// Wait for threads to complete
	for i := uint(0); i < h.threadCount; i++ {
		<-done
	}
}

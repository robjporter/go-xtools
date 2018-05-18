package xspinners

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type Spinner struct {
	out     io.Writer
	text    string
	running bool
	done    func()
	Name    string
	Delay   int
	Frames  []string
}

const erase = "\033[2K\r"

func New(spin Name, text string) *Spinner {
	s := get(spin)
	return &Spinner{out: os.Stdout, text: text, Name: s.Name, Delay: s.Delay, Frames: s.Frames}
}

func (s *Spinner) Start() {
	if !s.running {
		ctx, done := context.WithCancel(context.Background())
		t := time.NewTicker(time.Duration(s.Delay) * time.Millisecond)
		s.done = done
		s.running = true
		go func() {
			at := 0
			for {
				select {
				case <-ctx.Done():
					t.Stop()
					break
				case <-t.C:
					txt := erase + s.Frames[at%len(s.Frames)] + s.text
					fmt.Fprint(s.out, txt)
					at++
				}
			}
		}()
	}
}

func (s *Spinner) Stop() {
	if s.done != nil {
		s.done()
	}
	s.running = false
}

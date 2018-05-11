package xsignals

import (
	"errors"
	"os"
	"os/signal"
)

type Signal struct {
	sigchan  chan os.Signal
	termchan chan int
	Bindings map[os.Signal]func()
}

func (s *Signal) Bind(sig os.Signal, f func()) (err error) {
	for t, _ := range s.Bindings {
		if sig == t {
			err = errors.New("Signal is already bound to a function")
			return
		}
	}
	s.Bindings[sig] = f
	return nil
}

func New() *Signal {
	return &Signal{
		sigchan:  make(chan os.Signal),
		termchan: make(chan int),
		Bindings: make(map[os.Signal]func())}
}

func (s *Signal) Start() *Signal {
	go s.start()
	return s
}

func (s *Signal) start() {
	sigs := make([]os.Signal, len(s.Bindings))
	for t, _ := range s.Bindings {
		sigs = append(sigs, t)
	}
	signal.Notify(s.sigchan, sigs...)
	select {
	case sig := <-s.sigchan:
		s.Bindings[sig]()
	case _ = <-s.termchan:
		break
	}
}

// Package log provides a multi-tiered logger used to log output to a io.Writer.
// The different tiers are (least to highest piority):
//  1. Info
//  2. Debug
//  3. Warn
//  4. Error
//  5. Fatal
// One can set the output, output format, output time format, tags, and colors for tags.
// One can also add more tiers, and create their own logger with all the different tiers wanted.
package log

import (
	"io"
	"sync"
)

type Logger struct {
	tiers []ITier
	lock  sync.RWMutex
}

func NewLogger(tiers ...ITier) *Logger {
	l := &Logger{
		lock:  sync.RWMutex{},
		tiers: []ITier{},
	}
	return l.Add(tiers...)
}

// Add will add a tier to the tail end of the logger
func (l *Logger) Add(tiers ...ITier) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.tiers = append(l.tiers, tiers...)
	return l
}

// Get will get the tier at index idx, it will return nil if the
// index is out of bounds.
func (l *Logger) Get(idx int) ITier {
	if idx < len(l.tiers) {
		l.lock.RLock()
		defer l.lock.RUnlock()
		return l.tiers[idx]
	}
	return nil
}

// SetOutput sets where the logger will write to
func (l *Logger) SetOutput(out io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, tier := range l.tiers {
		tier.SetWriter(out)
	}
}

func (l *Logger) GetOutput() io.Writer {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if len(l.tiers) > 0 {
		return l.tiers[0].GetWriter()
	}
	return nil
}

// SetTimeFormat sets the time format for the time stamp on a log line
// Uses the go standard library timeformat format.
func (l *Logger) SetTimeFormat(format string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, tier := range l.tiers {
		tier.SetTimeFormat(NewTimeFormat(format))
	}
}

// SetFormat will set the logger to format all output.
// The format string
// MUST have a {TIME}, {TAG}, {MSG} string inside.
// For example: `{TIME} [{TAG}]:> {MSG}` will print logs of the form
// `10-21-1975 13:24:56 ERROR:> this is the message`
// Be careful, this will just do nothing if the format is invalid.
func (l *Logger) SetFormat(format string) {
	f := NewFormat(format)
	if f == nil {
		return
	}

	l.lock.Lock()
	defer l.lock.Unlock()
	for _, tier := range l.tiers {
		tier.SetFormat(f)
	}
}

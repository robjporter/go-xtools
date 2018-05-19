package log

import "sync"

// TimeFormat is the type used for formatting the time. Uses the same time format as
// the go standard library "time."
// example `Mon Jan _2 15:04:05 2006`
type TimeFormat struct {
	StringGetSetter
}

// NewTimeFormat creates a new TimeFormat
func NewTimeFormat(format string) ITimeFormat {
	return &TimeFormat{
		StringGetSetter{
			str:  format,
			lock: sync.RWMutex{},
		},
	}
}

// Dup duplicates the TimeFormat
func (f *TimeFormat) Dup() ITimeFormat {
	return NewTimeFormat(f.Get())
}

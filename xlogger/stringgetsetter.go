package log

import "sync"

// StringGetSetter is the base type for Color, Tag, Format, and TimeFormat
type StringGetSetter struct {
	str  string
	lock sync.RWMutex
}

// Get will return the string (thread safe)
func (s *StringGetSetter) Get() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.str
}

// Set will set the string (thread safe)
func (s *StringGetSetter) Set(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.str = str
}

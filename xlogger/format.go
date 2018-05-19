package log

import (
	"strings"
	"sync"
)

// Format struct holds the format of log messages. The Format string
// MUST have a {TIME}, {TAG}, {MSG} string inside.
// For example: `{TIME} [{TAG}]:> {MSG}` will print logs of the form
// `10-21-1975 13:24:56 ERROR:> this is the message`
type Format struct {
	StringGetSetter
}

// NewFormat creates a Format from a string.  If the format is not valid
// it will return nil
func NewFormat(format string) IFormat {
	if !validateFormat(format) {
		return nil
	}
	return &Format{
		StringGetSetter{
			str:  format,
			lock: sync.RWMutex{},
		},
	}
}

// Dup duplicates the format
func (f *Format) Dup() IFormat {
	return NewFormat(f.Get())
}

// Make sure the string contains all the keywords: {TIME}; {TAG}; {MSG};
func validateFormat(format string) bool {
	return strings.Contains(format, "{TIME}") &&
		strings.Contains(format, "{TAG}") &&
		strings.Contains(format, "{MSG}")
}

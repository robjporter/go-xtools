package log

import (
	"sync"
)

// Tag is the type used for Tags
type Tag struct {
	StringGetSetter
}

// NewTag creates a new Tag
func NewTag(tag string) ITag {
	return &Tag{
		StringGetSetter{
			str:  tag,
			lock: sync.RWMutex{},
		},
	}
}

// Dup duplicates the Tag
func (t *Tag) Dup() ITag {
	return NewTag(t.Get())
}

package log

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// Tier holds the properties for a logger tier.
type Tier struct {
	color      IColor
	tag        ITag
	format     IFormat
	timeFormat ITimeFormat
	writer     io.Writer
	lock       sync.Mutex
}

// NewTier Creates a new Tier with the desired properties
func NewTier(color IColor, tag ITag, format IFormat, timeFormat ITimeFormat, writer io.Writer) ITier {
	return &Tier{
		color:      color,
		tag:        tag,
		format:     format,
		timeFormat: timeFormat,
		writer:     writer,
		lock:       sync.Mutex{},
	}
}

// Dup duplicates a tier
func (t *Tier) Dup() ITier {
	return NewTier(t.GetColor(), t.GetTag(), t.GetFormat(), t.GetTimeFormat(), t.GetWriter())
}

// GetColor Returns the color property of the tier
func (t *Tier) GetColor() IColor {
	return t.color.Get()
}

// SetColor sets the color property of the tier
func (t *Tier) SetColor(c IColor) {
	t.color.Set(c.Get())
}

// GetTag returns the tag property
func (t *Tier) GetTag() ITag {
	return t.tag.Dup()
}

// SetTag sets the tag property
func (t *Tier) SetTag(tag ITag) {
	t.tag.Set(tag.Get())
}

// GetFormat gets the format property.
func (t *Tier) GetFormat() IFormat {
	return t.format.Dup()
}

// SetFormat sets the format property
func (t *Tier) SetFormat(format IFormat) {
	t.format.Set(format.Get())
}

//GetTimeFormat gets the time format property
func (t *Tier) GetTimeFormat() ITimeFormat {
	return t.timeFormat.Dup()
}

//SetTimeFormat sets the time format property
func (t *Tier) SetTimeFormat(timeFormat ITimeFormat) {
	t.timeFormat.Set(timeFormat.Get())
}

// GetWriter gets the writer property
func (t *Tier) GetWriter() io.Writer {
	return t.writer
}

// SetWriter sets the writer property
func (t *Tier) SetWriter(writer io.Writer) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.writer = writer
}

// Log will log a line to writer
func (t *Tier) Log(msgs ...interface{}) {
	out := ""
	for _, msg := range msgs {
		out = fmt.Sprintf("%s %v", out, msg)
	}
	out = out[1:]
	t.Logf(out)
}

// Logf will log a formated line to writer
func (t *Tier) Logf(format string, params ...interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	out := t.GetColor().Color(
		strings.Replace(strings.Replace(strings.Replace(t.GetFormat().Get(),
			"{TAG}", t.GetTag().Get(), -1),
			"{TIME}", time.Now().Format(t.GetTimeFormat().Get()), -1),
			"{MSG}", format, -1),
	)
	if len(params) > 0 {
		out = fmt.Sprintf(out, params...)
	}

	fmt.Fprintln(t.GetWriter(), out)
}

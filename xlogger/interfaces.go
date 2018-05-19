package log

import (
	"io"

	"github.com/popmedic/go-color/colorize"
)

type IStringGetter interface {
	Get() string
}

type IStringSetter interface {
	Set(string)
}

type IStringGetSetter interface {
	IStringGetter
	IStringSetter
}

type IColor interface {
	colorize.IColorize
	Get() IColor
	Set(IColor)
	Dup() IColor
}

type ITag interface {
	IStringGetSetter
	Dup() ITag
}

type IFormat interface {
	IStringGetSetter
	Dup() IFormat
}

type ITimeFormat interface {
	IStringGetSetter
	Dup() ITimeFormat
}

type ITier interface {
	GetColor() IColor
	SetColor(IColor)
	GetTag() ITag
	SetTag(ITag)
	GetFormat() IFormat
	SetFormat(IFormat)
	GetTimeFormat() ITimeFormat
	SetTimeFormat(ITimeFormat)
	GetWriter() io.Writer
	SetWriter(io.Writer)
	Dup() ITier
	Log(msg ...interface{})
	Logf(format string, params ...interface{})
}

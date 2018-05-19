package log

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/popmedic/go-color/colorize"
)

func TestNewTier(t *testing.T) {
	expTag := NewTag("TAG")
	expFormat := NewFormat("{TIME} {TAG}: {MSG}")
	expTimeFormat := NewTimeFormat("expTimeFormatStr")
	expWriter := os.Stdout

	res := NewTier(NewColor(colorize.NewColorize("yellow", "none")), expTag, expFormat, expTimeFormat, expWriter)
	if res == nil {
		t.Error("NewTier should not return nil")
	}
}

func TestGetSetTier(t *testing.T) {
	color := NewColor(colorize.NewColorize("black", "blue"))
	tag := NewTag("TAG")
	format := NewFormat("{TIME} {TAG}: {MSG}")
	timeFormat := NewTimeFormat("pqs")
	expColor := NewColor(colorize.NewColorize("yellow", "red"), colorize.NewColorize("", "").Add(NewColor()))
	expTag := NewTag("GAG")
	expFormat := NewFormat("{TIME} {TAG}: {MSG}")
	expTimeFormat := NewTimeFormat("2006-01-02 15:04:05")
	expWriter := os.Stdout

	tier := NewTier(color, tag, format, timeFormat, os.Stderr)
	tier.SetColor(expColor)
	tier.SetTag(expTag)
	tier.SetFormat(expFormat)
	tier.SetTimeFormat(expTimeFormat)
	tier.SetWriter(expWriter)

	if res := tier.Dup(); !compareTier(res, tier) {
		t.Errorf("given %+v expected %+v got %+v", tier, tier, res)
	}
}

func TestTierLog(t *testing.T) {
	expColor := NewColor(colorize.NewColorize("yellow", "none"))
	expTag := NewTag("GAG")
	expFormat := NewFormat("{TIME} : {TAG} : {MSG}")
	expTimeFormat := NewTimeFormat("2006-01-02 15:04:05")
	expWriter := bytes.NewBuffer([]byte{})
	tier := NewTier(expColor, expTag, expFormat, expTimeFormat, expWriter)

	tier.Logf("%s is a test %d", "This", 1)

	res := string(expWriter.Bytes())
	if !strings.HasSuffix(res, " : GAG : This is a test 1none\n") &&
		!strings.HasPrefix(res, "yellow") {
		t.Errorf("Logf: given \"TAG\", \"This is a test 1\\n\x1b[0m\" expected \"This is a test 1\\n\x1b[0m\" got %q", res)
	}

	expWriter.Reset()
	tier.Log("This is", "a", "test 1")
	res = string(expWriter.Bytes())
	if !strings.HasSuffix(res, " : GAG : This is a test 1none\n") &&
		!strings.HasPrefix(res, "yellow") {
		t.Errorf("Log: given \"TAG\", \"This is a test 1\\n\x1b[0m\" expected \"This is a test 1\\n\x1b[0m\" got %q", res)
	}
}

func compareColor(c1, c2 IColor) bool {
	return c1.Start() == c2.Start() &&
		c1.End() == c2.End()
}

func compareTier(t1, t2 ITier) bool {
	return compareColor(t1.GetColor(), t2.GetColor()) &&
		compareTag(t1.GetTag(), t2.GetTag()) &&
		compareFormat(t1.GetFormat(), t2.GetFormat()) &&
		compareTimeFormat(t1.GetTimeFormat(), t2.GetTimeFormat()) &&
		t1.GetWriter() == t2.GetWriter()
}

func compareFormat(f1, f2 IFormat) bool {
	return f1.Get() == f2.Get()
}

func compareTag(t1, t2 ITag) bool {
	return t1.Get() == t2.Get()
}

func compareTimeFormat(f1, f2 ITimeFormat) bool {
	return f1.Get() == f2.Get()
}

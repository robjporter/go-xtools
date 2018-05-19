package log

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/popmedic/go-color/colorize/html"
	"github.com/popmedic/go-color/colorize/tty"
	"github.com/popmedic/go-writers/linewriter"
)

const defaultMaxBufferLines = 20

const (
	infoIdx int = iota
	debugIdx
	warnIdx
	errorIdx
	fatalIdx
)

const (
	defaultFormatString     = "{TIME} [{TAG}]> {MSG}"
	defaultTimeFormatString = "2006-01-02 15:04:05"
)

const (
	defaultInfoTagString  = "INFO"
	defaultDebugTagString = "DEBUG"
	defaultWarnTagString  = "WARN"
	defaultErrorTagString = "ERROR"
	defaultFatalTagString = "FATAL"
)

var (
	defaultOut     = os.Stdout
	defaultHTMLOut = linewriter.NewLineWriter(defaultMaxBufferLines)
)

var (
	defaultFormat     = NewFormat(defaultFormatString)
	defaultTimeFormat = NewTimeFormat(defaultTimeFormatString)
)

var (
	defaultInfoColor  = NewColor(tty.Reset())
	defaultDebugColor = NewColor(tty.FgGreen())
	defaultWarnColor  = NewColor(tty.FgYellow())
	defaultErrorColor = NewColor(tty.FgRed())
	defaultFatalColor = NewColor(tty.BgRed().Add(tty.FgHiWhite()))

	defaultInfoHTMLColor  = NewColor(html.FgWhite(), html.BgBlack(), html.Monospace())
	defaultDebugHTMLColor = NewColor(html.FgGreen(), html.BgBlack(), html.Monospace())
	defaultWarnHTMLColor  = NewColor(html.FgYellow(), html.BgBlack(), html.Monospace())
	defaultErrorHTMLColor = NewColor(html.FgRed(), html.BgBlack(), html.Monospace())
	defaultFatalHTMLColor = NewColor(html.BgRed(), html.FgWhite(), html.Bold(), html.Monospace())
)

var (
	defaultInfoTag  = NewTag(defaultInfoTagString)
	defaultDebugTag = NewTag(defaultDebugTagString)
	defaultWarnTag  = NewTag(defaultWarnTagString)
	defaultErrorTag = NewTag(defaultErrorTagString)
	defaultFatalTag = NewTag(defaultFatalTagString)
)

var (
	infoTier = NewTier(
		defaultInfoColor,
		defaultInfoTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	debugTier = NewTier(
		defaultDebugColor,
		defaultDebugTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	warnTier = NewTier(
		defaultWarnColor,
		defaultWarnTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	errorTier = NewTier(
		defaultErrorColor,
		defaultErrorTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)
	fatalTier = NewTier(
		defaultFatalColor,
		defaultFatalTag,
		defaultFormat,
		defaultTimeFormat,
		defaultOut,
	)

	infoHTMLTier = NewTier(
		defaultInfoHTMLColor,
		defaultInfoTag,
		defaultFormat,
		defaultTimeFormat,
		defaultHTMLOut,
	)
	debugHTMLTier = NewTier(
		defaultDebugHTMLColor,
		defaultDebugTag,
		defaultFormat,
		defaultTimeFormat,
		defaultHTMLOut,
	)
	warnHTMLTier = NewTier(
		defaultWarnHTMLColor,
		defaultWarnTag,
		defaultFormat,
		defaultTimeFormat,
		defaultHTMLOut,
	)
	errorHTMLTier = NewTier(
		defaultErrorHTMLColor,
		defaultErrorTag,
		defaultFormat,
		defaultTimeFormat,
		defaultHTMLOut,
	)
	fatalHTMLTier = NewTier(
		defaultFatalHTMLColor,
		defaultFatalTag,
		defaultFormat,
		defaultTimeFormat,
		defaultHTMLOut,
	)
)

var (
	tiers = NewLogger(
		infoTier,
		debugTier,
		warnTier,
		errorTier,
		fatalTier,
	)
	htmlTiers = NewLogger(
		infoHTMLTier,
		debugHTMLTier,
		warnHTMLTier,
		errorHTMLTier,
		fatalHTMLTier,
	)

	htmlStatus     = false
	htmlCloser     io.Closer
	htmlStatusLock = sync.RWMutex{}
)

// SetOutput sets where the logger will write to
func SetOutput(out io.Writer) {
	tiers.SetOutput(out)
}

// SetTimeFormat sets the time format for the time stamp on a log line
// Uses the go standard library timeformat format.
func SetTimeFormat(format string) {
	tiers.SetTimeFormat(format)
}

// SetFormat will set the logger to format all output.
// The format string
// MUST have a {TIME}, {TAG}, {MSG} string inside.
// For example: `{TIME} [{TAG}]:> {MSG}` will print logs of the form
// `10-21-1975 13:24:56 ERROR:> this is the message`
// Be careful, this will just do nothing if the format is invalid.
func SetFormat(format string) {
	tiers.SetFormat(format)
}

func SetHTMLStatus(status bool, addr string) {
	htmlStatusLock.Lock()
	defer htmlStatusLock.Unlock()
	htmlStatus = status
	if status {
		var err error
		htmlCloser, err = listenAndServeWithClose(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(htmlLoggerString()))
		}))
		if nil != err {
			log(errorIdx, "UNABLE TO START HTML SERVER: ", err)
			htmlCloser, htmlStatus = nil, false
		}
	} else if nil != htmlCloser {
		if err := htmlCloser.Close(); nil != err {
			log(errorIdx, "UNABLE TO TERMINATE HTML SERVER: ", err)
			htmlCloser, htmlStatus = nil, false
		}
	}
}

func GetHTMLStatus() bool {
	htmlStatusLock.RLock()
	defer htmlStatusLock.RUnlock()
	return htmlStatus
}

// SetHTMLOutput sets where the logger will write to
func SetHTMLOutput(out io.Writer) {
	htmlTiers.SetOutput(out)
}

func GetHTMLOutput() io.Writer {
	return htmlTiers.GetOutput()
}

// Get will get the tier at idx
func Get(idx int) ITier {
	return tiers.Get(idx)
}

// GetInfo gets the info tier
func GetInfo() ITier {
	return Get(infoIdx)
}

// GetDebug gets the info tier
func GetDebug() ITier {
	return Get(debugIdx)
}

// GetWarn gets the info tier
func GetWarn() ITier {
	return Get(warnIdx)
}

// GetError gets the info tier
func GetError() ITier {
	return Get(errorIdx)
}

// GetFatal gets the info tier
func GetFatal() ITier {
	return Get(fatalIdx)
}

// Infof prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Info tag and color.
func Infof(f string, i ...interface{}) {
	logf(infoIdx, f, i...)
}

// Info prints to output values i joined with a space.
// Will add the assigned Info tag and color.
func Info(i ...interface{}) {
	log(infoIdx, i...)
}

// Debugf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Debug tag and color.
func Debugf(f string, i ...interface{}) {
	logf(debugIdx, f, i...)
}

// Debug prints to output values i joined with a space.
// Will add the assigned Debug tag and color.
func Debug(i ...interface{}) {
	log(debugIdx, i...)
}

// Warnf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Warn tag and color.
func Warnf(f string, i ...interface{}) {
	logf(warnIdx, f, i...)
}

// Warn prints to output values i joined with a space.
// Will add the assigned Warn tag and color.
func Warn(i ...interface{}) {
	log(warnIdx, i...)
}

// Errorf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Error tag and color.
func Errorf(f string, i ...interface{}) {
	logf(errorIdx, f, i...)
}

// Error prints to output values i joined with a space.
// Will add the assigned Error tag and color.
func Error(i ...interface{}) {
	log(errorIdx, i...)
}

// Fatalf prints to output f formatted with values i.
// Uses the go standard library style format strings.
// Will add the assigned Fatal tag and color.
func Fatalf(exit func(int), f string, i ...interface{}) {
	logf(fatalIdx, f, i...)
	exit(1)
}

// Fatal prints to output values i joined with a space.
// Will add the assigned Fatal tag and color.
func Fatal(exit func(int), i ...interface{}) {
	log(fatalIdx, i...)
	exit(1)
}

func htmlLoggerString() string {
	o := `<!DOCTYPE html>
<html style="background-color:grey;">
<head>
<meta http-equiv="refresh" content="5; URL=/here">
</head>
<body style="width:80%;margin-left:auto;margin-right:auto;background-color:Black;padding:12px;border-style:solid;border-color:white;border-width:1px;border-radius:4px">
<h1 style="font-family:'Consola', monospace;font-size:18pt;color:white;">Popmedia Log</h1>
`
	if v, ok := htmlTiers.GetOutput().(*linewriter.LineWriter); ok {
		for _, item := range v.All() {
			o = o + fmt.Sprint(item) + "<br/>\n"
		}
		for n := v.Max() - v.Len(); n > 0; n-- {
			o = o + "<br/>"
		}
	}
	return o + `</body>
</html>`
}

func log(idx int, i ...interface{}) {
	if tier := Get(idx); tier != nil {
		tier.Log(i...)
	}
	if GetHTMLStatus() {
		if tier := htmlTiers.Get(idx); tier != nil {
			tier.Log(i...)
		}
	}
}

func logf(idx int, f string, i ...interface{}) {
	if tier := Get(idx); tier != nil {
		tier.Logf(f, i...)
	}
	if GetHTMLStatus() {
		if tier := htmlTiers.Get(idx); tier != nil {
			tier.Logf(f, i...)
		}
	}
}

func listenAndServeWithClose(addr string, handler http.Handler) (io.Closer, error) {
	srv := &http.Server{Addr: addr, Handler: handler}

	if addr == "" {
		addr = ":http"
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		err := srv.Serve(listener)
		if err != nil {
			log(errorIdx, "UNABLE TO ACTIVATE HTML LOGGER: HTTP Server Error - ", err)
		}
	}()

	return listener, nil
}

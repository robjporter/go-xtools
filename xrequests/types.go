package xrequests

import (
	"net/http"
	"net/url"
	"time"
)

type RetryType struct {
	RetryStatus []int
	RetryTime   time.Duration
	RetryCount  int
	Attempt     int
	Enable      bool
}

type Agent struct {
	BasicAuth         struct{ Username, Password string }
	BounceToRawString bool
	ClearAgentData    bool
	ClearSendData     bool
	Client            *http.Client
	Cookies           []*http.Cookie
	Data              map[string]interface{}
	Debug             bool
	Errors            []error
	FileData          []File
	ForceType         string
	FormData          url.Values
	Header            http.Header
	Method            string
	QueryData         url.Values
	RawString         string
	Retryer           *RetryType
	SliceData         []interface{}
	TargetType        string
	Transport         *http.Transport
	URL               string
}

type File struct {
	Filename  string
	Fieldname string
	Data      []byte
}

type Request *http.Request
type Response *http.Response

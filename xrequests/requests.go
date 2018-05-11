package xrequests

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

var DisableTransportSwap bool

func New() *Agent {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, _ := cookiejar.New(&cookiejarOptions)

	a := &Agent{
		BasicAuth:         struct{ Username, Password string }{},
		BounceToRawString: false,
		ClearAgentData:    true,
		Client:            &http.Client{Jar: jar},
		Cookies:           make([]*http.Cookie, 0),
		Data:              make(map[string]interface{}),
		Debug:             false,
		Errors:            nil,
		FileData:          make([]File, 0),
		FormData:          url.Values{},
		Header:            http.Header{},
		QueryData:         url.Values{},
		RawString:         "",
		Retryer:           &RetryType{},
		SliceData:         []interface{}{},
		TargetType:        TypeJSON,
		Transport:         &http.Transport{},
	}

	DisableTransportSwap = false

	return a
}

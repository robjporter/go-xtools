package xrequests

import (
	"net/http"
	"net/url"
)

func (a *Agent) Proxy(proxyurl string) *Agent {
	parsed, err := url.Parse(proxyurl)
	if err != nil {
		a.Errors = append(a.Errors, err)
	} else if proxyurl == "" {
		a.Transport.Proxy = nil
	} else {
		a.Transport.Proxy = http.ProxyURL(parsed)
	}
	return a
}

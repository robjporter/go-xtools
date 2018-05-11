package xrequests

import (
	"errors"
	"net/http"
	"net/url"
)

func (a *Agent) SetDebug(enabled bool) *Agent {
	a.Debug = enabled
	return a
}

func (a *Agent) SetDoNotClearBetweenSessions() *Agent {
	a.ClearAgentData = false
	return a
}

func (a *Agent) SetClearSendDataBetweenSessions() *Agent {
	a.ClearSendData = true
	return a
}

func (a *Agent) ClearAgent() {
	a.Errors = nil
	if !a.ClearAgentData {
		return
	}
	a.BounceToRawString = false
	a.URL = ""
	a.Method = ""
	a.Header = http.Header{}
	a.Data = make(map[string]interface{})
	a.SliceData = []interface{}{}
	a.FormData = url.Values{}
	a.FileData = make([]File, 0)
	a.QueryData = url.Values{}
	a.RawString = ""
	a.ForceType = ""
	a.Cookies = make([]*http.Cookie, 0)
	a.Errors = nil
}

func (a *Agent) SetHeader(param, value string) *Agent {
	a.Header.Set(param, value)
	return a
}

func (a *Agent) AppendHeader(param, value string) *Agent {
	a.Header.Add(param, value)
	return a
}

func (a *Agent) SetBasicAuth(username, password string) *Agent {
	a.BasicAuth = struct{ Username, Password string }{username, password}
	return a
}

func (a *Agent) AddCookie(c *http.Cookie) *Agent {
	a.Cookies = append(a.Cookies, c)
	return a
}

func (a *Agent) AddCookies(c []*http.Cookie) *Agent {
	a.Cookies = append(a.Cookies, c...)
	return a
}

func (a *Agent) Type(str string) *Agent {
	if _, ok := Types[str]; ok {
		a.ForceType = str
	} else {
		a.Errors = append(a.Errors, errors.New("Incorrect Type: "+str))
	}
	return a
}

func (a *Agent) Param(key, value string) *Agent {
	a.QueryData.Add(key, value)
	return a
}

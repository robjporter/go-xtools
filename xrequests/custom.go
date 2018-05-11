package xrequests

import (
	"crypto/tls"
	"net/url"
)

func (a *Agent) AddBearer(token string) *Agent {
	a.SetHeader("Authorization", "Bearer "+token)
	return a
}

func (a *Agent) IgnoreTLSCheck() *Agent {
	a.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return a
}

func (a *Agent) SendRawBody() *Agent {
	a.BounceToRawString = true
	return a
}

func (a *Agent) ClearSendDataNow() *Agent {
	if !a.ClearSendData {
		return a
	}
	a.RawString = ""
	a.Data = make(map[string]interface{})
	a.SliceData = []interface{}{}
	a.FormData = url.Values{}
	a.FileData = make([]File, 0)
	return a
}

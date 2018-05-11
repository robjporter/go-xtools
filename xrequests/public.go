package xrequests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (a *Agent) End(callback ...func(response Response, body string, errs []error)) (Response, string, []error) {
	var bytesCallback []func(response Response, body []byte, errs []error)

	if len(callback) > 0 {
		bytesCallback = []func(response Response, body []byte, errs []error){
			func(response Response, body []byte, errs []error) {
				callback[0](response, string(body), errs)
			},
		}
	}

	resp, body, errs := a.EndBytes(bytesCallback...)
	bodyString := string(body)
	return resp, bodyString, errs
}

func (a *Agent) EndBytes(callback ...func(response Response, body []byte, errs []error)) (Response, []byte, []error) {
	var errs []error
	var resp Response
	var body []byte

	for {
		resp, body, errs = a.getResponseBytes()
		if errs != nil {
			return nil, nil, errs
		}
		if a.isRetryableRequest(resp) {
			resp.Header.Set("Retry-Count", strconv.Itoa(a.Retryer.Attempt))
			break
		}
	}

	respCallback := *resp
	if len(callback) != 0 {
		callback[0](&respCallback, body, a.Errors)
	}

	return resp, body, nil
}

func (a *Agent) MakeRequest() (*http.Request, error) {
	var req *http.Request
	var contentType string
	var contentReader io.Reader
	var err error

	if a.Method == "" {
		return nil, errors.New("No Method specified")
	}

	switch a.TargetType {
	case TypeJSON:
		contentType, contentReader = a.processTypeJSON()
	case TypeForm, TypeFormData, TypeUrlencoded:
		contentType, contentReader = a.processTypeForm()
	case TypeText:
		contentType, contentReader = a.processTypeText()
	case TypeXML:
		contentType, contentReader = a.processTypeXML()
	case TypeMultipart:
		contentType, contentReader = a.processTypeMultipart()
	default:
		return nil, errors.New("TargetType `" + a.TargetType + "` could not be determined.")
	}

	if req, err = http.NewRequest(a.Method, a.URL, contentReader); err != nil {
		return nil, err
	}

	for k, v := range a.Header {
		for _, v2 := range v {
			req.Header.Add(k, v2)
		}
		if strings.EqualFold(k, "Host") {
			req.Host = v[0]
		}
	}

	if len(contentType) != 0 && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentType)
	}

	q := req.URL.Query()
	for k, v := range a.QueryData {
		for _, v2 := range v {
			q.Add(k, v2)
		}
	}

	if a.BasicAuth != struct{ Username, Password string }{} {
		req.SetBasicAuth(a.BasicAuth.Username, a.BasicAuth.Password)
	}

	for _, cookie := range a.Cookies {
		req.AddCookie(cookie)
	}

	return req, nil
}

func (a *Agent) EndStruct(v interface{}, callback ...func(response Response, v interface{}, body []byte, errs []error)) (Response, []byte, []error) {
	resp, body, errs := a.EndBytes()

	if errs != nil {
		return nil, body, errs
	}

	err := json.Unmarshal(body, &v)
	if err != nil {
		a.Errors = append(a.Errors, err)
		return resp, body, a.Errors
	}

	respCallback := *resp
	if len(callback) != 0 {
		callback[0](&respCallback, v, body, a.Errors)
	}

	return resp, body, nil
}

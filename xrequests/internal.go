package xrequests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

func (a *Agent) isRetryableRequest(resp Response) bool {
	if a.Retryer.Enable && a.Retryer.Attempt < a.Retryer.RetryCount && contains(resp.StatusCode, a.Retryer.RetryStatus) {
		time.Sleep(a.Retryer.RetryTime)
		a.Retryer.Attempt++
		return false
	}
	return true
}

func (a *Agent) getResponseBytes() (Response, []byte, []error) {
	var req *http.Request
	var err error
	var resp Response

	if len(a.Errors) != 0 {
		return nil, nil, a.Errors
	}

	switch a.ForceType {
	case TypeJSON, TypeForm, TypeXML, TypeText, TypeMultipart:
		a.TargetType = a.ForceType
	default:
		contentType := a.Header.Get("Content-Type")
		for k, v := range Types {
			if contentType == v {
				a.TargetType = k
			}
		}
	}

	if len(a.Data) != 0 && len(a.SliceData) != 0 {
		a.BounceToRawString = true
	}

	req, err = a.MakeRequest()
	if err != nil {
		a.Errors = append(a.Errors, err)
		return nil, nil, a.Errors
	}

	if !DisableTransportSwap {
		a.Client.Transport = a.Transport
	}

	if a.Debug {
		dump, err := httputil.DumpRequest(req, true)
		fmt.Println("=================================================")
		fmt.Println(string(dump))
		fmt.Println(err)
		fmt.Println("=================================================")
		//LOGGER
		/*
			s.logger.SetPrefix("[http] ")
					if err != nil {
						s.logger.Println("Error:", err)
					} else {
						s.logger.Printf("HTTP Request: %s", string(dump))
					}
		*/
	}

	resp, err = a.Client.Do(req)
	if err != nil {
		a.Errors = append(a.Errors, err)
		return nil, nil, a.Errors
	}
	defer resp.Body.Close()

	if a.Debug {
		dump, err := httputil.DumpRequest(req, true)
		fmt.Println("=================================================")
		fmt.Println(string(dump))
		fmt.Println(err)
		fmt.Println("=================================================")
		//LOGGER
		/*
			s.logger.SetPrefix("[http] ")
					if err != nil {
						s.logger.Println("Error:", err)
					} else {
						s.logger.Printf("HTTP Request: %s", string(dump))
					}
		*/
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return resp, body, nil
}

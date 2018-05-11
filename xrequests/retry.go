package xrequests

import (
	"errors"
	"net"
	"net/http"
	"strconv"
	"time"
)

func (a *Agent) Retry(count int, rTime time.Duration, codes ...int) *Agent {
	for _, code := range codes {
		text := http.StatusText(code)
		if len(text) == 0 {
			a.Errors = append(a.Errors, errors.New("StatusCode `"+strconv.Itoa(code)+"` does not exist"))
		}
	}

	a.Retryer = &RetryType{RetryStatus: codes, RetryTime: rTime, RetryCount: count, Attempt: 0, Enable: true}
	return a
}

func (a *Agent) Timeout(timeout time.Duration) *Agent {
	a.Transport.Dial = func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, timeout)
		if err != nil {
			a.Errors = append(a.Errors, err)
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(timeout))
		return conn, nil
	}
	return a
}

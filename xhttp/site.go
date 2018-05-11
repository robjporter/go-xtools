package xhttp

import (
	"fmt"
	"net/http"
)

type Status int

const (
	UNCHECKED Status = iota
	DOWN
	UP
)

// The Site struct encapsulates the details about the site being monitored.
type Site struct {
	Url                string
	Last_status        Status
	Last_status_string string
}

func NewSite() *Site {
	return &Site{}
}

func (s *Site) SetURL(url string) {
	s.Url = url
}

func (s *Site) LastStatus() int {
	return int(s.Last_status)
}

func (s *Site) LastStatusString() string {
	return s.Last_status_string
}

// Site.Status makes a GET request to a given URL and checks whether or not the
// resulting status code is 200.
func (s *Site) Status() (Status, error) {
	fmt.Println(s)
	resp, err := http.Get(s.Url)
	status := s.Last_status

	if (err == nil) && (resp.StatusCode == 200) {
		status = UP
	} else {
		status = DOWN
	}

	s.Last_status = status
	s.Last_status_string = lastStatusString(s.Last_status)

	return status, err
}

func lastStatusString(i Status) string {
	if i == UNCHECKED {
		return "UNCHECKED"
	} else if i == DOWN {
		return "DOWN"
	} else if i == UP {
		return "UP"
	}
	return ""
}

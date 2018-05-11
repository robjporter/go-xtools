package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"../xrequests"
)

func main() {
	a := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10*time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
	response, body, err := a.Post("https://10.51.31.147/nuova").Send(`<aaaLogin inName="admin" inPassword="C15co123" />`).End()
	if err == nil {
		fmt.Printf("STATUS CODE: %d\n", response.StatusCode)
		fmt.Printf("BODY: %s\n", body)
	}
	var login Login
	xml.Unmarshal([]byte(body), &login)

	response, body, err = a.Send(`<aaaLogout inCookie="` + login.OutCookie + `" />`).End()
	if err == nil {
		fmt.Printf("STATUS CODE: %d\n", response.StatusCode)
		fmt.Printf("BODY: %s\n", body)
	}
}

type Login struct {
	XMLName          xml.Name `xml:"aaaLogin"`
	Cookie           string   `xml:"cookie,attr"`
	Response         string   `xml:"response,attr"`
	OutCookie        string   `xml:"outCookie,attr"`
	OutRefreshPeriod string   `xml:"outRefreshPeriod,attr"`
	OutPriv          string   `xml:"outPriv,attr"`
	OutDomains       string   `xml:"outDomains,attr"`
	OutChannel       string   `xml:"outChannel,attr"`
	OutEvtChannel    string   `xml:"outEvtChannel,attr"`
	OutSessionID     string   `xml:"outSessionId,attr"`
	OutVersion       string   `xml:"outVersion,attr"`
	OutName          string   `xml:"outName,attr"`
}

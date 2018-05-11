package aci

import (
	"encoding/json"
	"strings"

	"github.com/robjporter/go-utils/go/as"
	"github.com/robjporter/go-utils/web/request"
)

const (
	JSON_REPLACEMENT_START = "<"
	JSON_REPLACEMENT_END   = ">"
	JSON_LOGIN             = `{"aaaUser": {"attributes": {"name": "<USERNAME>", "pwd": "<PASSWORD>"}}}`
	JSON_LOGOUT            = `{"aaaUser":{"attributes":{"name": "<USERNAME>"}}}`
)

type ACIDATA struct {
	version   string
	token     string
	privRead  string
	privWrite string
}

type ACILogin struct {
	ip       string
	username string
	password string
}

type RESPONSE struct {
	Response string
	Body     string
	Errors   []error
}

type ACI struct {
	handler      *request.SuperAgent
	cookie       string
	status       bool
	replacements map[string]string
	data         ACIDATA
	login        ACILogin
	LastResponse RESPONSE
}

func New() *ACI {
	a := ACI{
		handler: request.New(),
		login: ACILogin{
			ip:       "",
			username: "",
			password: "",
		},
		cookie:       "",
		status:       false,
		replacements: make(map[string]string),
	}
	a.handler.SetInsecureDefaults()
	a.addReplacementDefaults()
	a.handler.SetRecorder(false)
	return &a
}

//PRIVATE*********************************************************************

func (a *ACI) internalLogin() {
	var resp2, body2 string
	var err2 []error
	a.addReplacementString("USERNAME", a.login.username)
	a.addReplacementString("PASSWORD", a.login.password)
	json := a.jsonReplace(JSON_LOGIN)
	resp, body, err := a.handler.Post("https://"+a.login.ip+"/api/aaaLogin.json").Set("Content-Type", "application/json").Send(json).End()
	if err == nil {
		if resp.StatusCode == 200 {
			if a.getCookieVersion(body) {
				a.status = true
				resp2 = as.ToString(resp)
				body2 = as.ToString(body)
			}
		}
	}
	err2 = err
	a.LastResponse.Response = resp2
	a.LastResponse.Body = body2
	a.LastResponse.Errors = err2
}

func (a *ACI) getCookieVersion(json string) bool {
	a.data.token = getJSONData(json, "imdata[0].aaaLogin.attributes.token")
	a.data.version = getJSONData(json, "imdata[0].aaaLogin.attributes.version")
	a.data.privRead = getJSONData(json, "imdata[0].aaaLogin.children[0].aaaUserDomain.attributes.rolesR")
	a.data.privWrite = getJSONData(json, "imdata[0].aaaLogin.children[0].aaaUserDomain.attributes.rolesW")
	if a.data.token != "unknown" && a.data.version != "unknown" {
		a.addReplacementString("TOKEN", a.data.token)
		return true
	}
	return false
}

func (a *ACI) jsonReplace(json string) string {
	for k, v := range a.replacements {
		k = JSON_REPLACEMENT_START + k + JSON_REPLACEMENT_END
		json = strings.Replace(json, k, v, -1)
	}
	return json
}

func (a *ACI) addReplacementDefaults() {

}

func (a *ACI) addReplacementString(key, value string) {
	a.replacements[key] = value
}

func getJSONData(jsonstr string, element string) string {
	output := ""
	if jsonstr != "" {
		var data2 interface{}
		json.Unmarshal([]byte(jsonstr), &data2)

		tmp, err := "", "" //jmespath.Search(element, data2)
		if err == "" {
			output = as.ToString(tmp)
		}
	}
	return output
}

//PUBLIC***********************************************************************

func (a *ACI) Login(ip, username, password string) *ACI {
	a.login.ip = ip
	a.login.username = username
	a.login.password = password
	a.internalLogin()
	return a
}

func (a *ACI) Logout() []error {
	json := a.jsonReplace(JSON_LOGOUT)
	resp, _, err := a.handler.Post("https://"+a.login.ip+"/api/aaaLogout.json").Set("Content-Type", "application/json").Send(json).End()
	if err == nil {
		if resp.StatusCode == 200 {
			a.status = false
			return nil
		}
	}
	return err
}

func (a *ACI) End() (string, []error) {
	a.Logout()
	return "FINSIHED", nil
}

func (a *ACI) GetPriviledges() string {
	output := ""
	if a.data.privRead != "" && a.data.privRead != "unknown" {
		output = a.data.privRead
		if a.data.privWrite != "" && a.data.privWrite != "unknown" {
			output += "/" + a.data.privWrite
		}
	}
	return output
}

func (a *ACI) GetVersion() string {
	if a.data.version != "" && a.data.version != "unknown" {
		return a.data.version
	}
	return ""
}

func (a *ACI) IsAdmin() bool {
	if a.data.privRead == "admin" && a.data.privWrite != "admin" {
		return true
	}
	return false
}

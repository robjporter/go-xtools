package ucs

import (
	//"fmt"

	"strings"

	"github.com/robjporter/go-utils/go/as"
	xmlx "github.com/robjporter/go-utils/go/encode/xml"
	"github.com/robjporter/go-utils/web/request"
)

const (
	XML_REPLACEMENT_START = "<"
	XML_REPLACEMENT_END   = ">"
	XML_LOGIN             = `<aaaLogin inName="<USERNAME>" inPassword="<PASSWORD>"/>`
	XML_LOGOUT            = `<aaaLogout inCookie="<COOKIE>"/>`
	XML_TOP_SYSTEM_INFO   = `<configResolveClass cookie="<COOKIE>" classId="topSystem" inHierarchical="false"/>`
	XML_SERVER_DN         = `<configFindDnsByClassId cookie="<COOKIE>" classId="computeItem"/>`
	XML_SERVER_DETAIL_DN  = `<configResolveDn cookie="<COOKIE>" dn="<DN>"  inHierarchical="false"/>`
)

type UCSMDATA struct {
	version string
	cookie  string
	priv    string
	name    string
}

type UCSMLogin struct {
	ip       string
	username string
	password string
}

type RESPONSE struct {
	Response string
	Body     string
	Errors   []error
}

type UCSM struct {
	handler      *request.SuperAgent
	cookie       string
	status       bool
	replacements map[string]string
	data         UCSMDATA
	login        UCSMLogin
	LastResponse RESPONSE
}

type UCSBladeInfo struct {
	BladeDescr        string
	BladeDN           string
	BladeLabel        string
	BladeModel        string
	BladeName         string
	BladePID          string
	BladeSlot         string
	BladeChassis      string
	BladeSerial       string
	BladeUUID         string
	BladeMemory       string
	BladeAssociation  string
	BladeAssociatedTo string
	BladeAvailability string
	BladeSockets      string
	BladeCores        string
	BladePower        string
}

func New() *UCSM {
	u := UCSM{
		handler: request.New(),
		login: UCSMLogin{
			ip:       "",
			username: "",
			password: "",
		},
		cookie:       "",
		status:       false,
		replacements: make(map[string]string),
	}
	u.handler.SetInsecureDefaults()
	u.addReplacementDefaults()
	u.handler.SetRecorder(false)
	return &u
}

//PRIVATE*********************************************************************

func (u *UCSM) internalLogin() {
	var resp2, body2 string
	var err2 []error
	u.addReplacementString("USERNAME", u.login.username)
	u.addReplacementString("PASSWORD", u.login.password)
	xml := u.xmlReplace(XML_LOGIN)
	resp, body, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			if u.getCookieVersion(body) {
				u.status = true
				resp2 = as.ToString(resp)
				body2 = as.ToString(body)
			}
		}
	}
	err2 = err
	u.LastResponse.Response = resp2
	u.LastResponse.Body = body2
	u.LastResponse.Errors = err2
}

func (u *UCSM) getCookieVersion(xml string) bool {
	u.data.cookie = getXMLAttributeData(xml, "aaaLogin", "", "outCookie")
	u.data.version = getXMLAttributeData(xml, "aaaLogin", "", "outVersion")
	u.data.priv = getXMLAttributeData(xml, "aaaLogin", "", "outPriv")
	if u.data.cookie != "unknown" && u.data.version != "unknown" {
		u.addReplacementString("COOKIE", u.data.cookie)
		return true
	}
	return false
}

func (u *UCSM) GetSystemName() string {
	u.data.name = ""
	u.addReplacementString("COOKIE", u.data.cookie)
	xml := u.xmlReplace(XML_TOP_SYSTEM_INFO)
	resp, body, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			u.data.name = getXMLAttributeData(body, "configResolveClass", "topSystem", "name")
		}
	}
	u.LastResponse.Response = as.ToString(resp)
	u.LastResponse.Body = body
	u.LastResponse.Errors = err
	return u.data.name
}

func (u *UCSM) xmlReplace(xml string) string {
	for k, v := range u.replacements {
		k = XML_REPLACEMENT_START + k + XML_REPLACEMENT_END
		xml = strings.Replace(xml, k, v, -1)
	}
	return xml
}

func (u *UCSM) addReplacementDefaults() {

}

func (u *UCSM) addReplacementString(key, value string) {
	u.replacements[key] = value
}

func getXMLAttributeData(xml string, root string, element string, attribute string) string {
	doc := xmlx.New()
	err := doc.LoadString(xml, nil)

	if err == nil {
		if element == "" {
			nod := doc.SelectNode("", root)
			if nod != nil {
				return nod.As("", attribute)
			}
		} else {
			nod := doc.SelectNode("", root)
			if nod != nil {
				nod2 := nod.SelectNode("", element)
				if nod2 != nil {
					return nod2.As("", attribute)
				}
			}
		}
	}

	return "unknown"
}

func getXMLElements(xml string, root string, element string) []*xmlx.Node {
	doc := xmlx.New()
	err := doc.LoadString(xml, nil)

	if err == nil {
		nod := doc.SelectNode("", root)
		if nod != nil {
			list := nod.SelectNodes("", "dn")
			if list != nil {
				return list
			}
		}
	}

	return []*xmlx.Node{}
}

//PUBLIC***********************************************************************

func (u *UCSM) Login(ip, username, password string) *UCSM {
	u.login.ip = ip
	u.login.username = username
	u.login.password = password
	u.internalLogin()
	return u
}

func (u *UCSM) Logout() []error {
	xml := u.xmlReplace(XML_LOGOUT)
	resp, _, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			u.status = false
			return nil
		}
	}
	return err
}

func (u *UCSM) End() (string, []error) {
	u.Logout()
	return "FINSIHED", nil
}

func (u *UCSM) GetPriviledges() string {
	if u.data.priv != "" && u.data.priv != "unknown" {
		return u.data.priv
	}
	return ""
}

func (u *UCSM) GetVersion() string {
	if u.data.version != "" && u.data.version != "unknown" {
		return u.data.version
	}
	return ""
}

func (u *UCSM) IsAdmin() bool {
	if u.data.priv != "" && u.data.priv != "unknown" {
		if strings.Contains(u.data.priv, "admin") {
			return true
		}
	}
	return false
}

func (u *UCSM) GetSystemBlades() []UCSBladeInfo {
	servers := []UCSBladeInfo{}
	u.addReplacementString("COOKIE", u.data.cookie)
	xml := u.xmlReplace(XML_SERVER_DN)
	resp, body, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			list := getXMLElements(body, "configFindDnsByClassId", "dn")
			for i := 0; i < len(list); i++ {
				servers = append(servers, u.getServerDetail(getXMLAttributeData(as.ToString(list[i]), "dn", "", "value")))
			}
		}
	}
	u.LastResponse.Response = as.ToString(resp)
	u.LastResponse.Body = body
	u.LastResponse.Errors = err
	return servers
}

func (u *UCSM) getServerDetail(dn string) UCSBladeInfo {
	server := UCSBladeInfo{}
	u.addReplacementString("COOKIE", u.data.cookie)
	u.addReplacementString("DN", dn)
	xml := u.xmlReplace(XML_SERVER_DETAIL_DN)
	resp, body, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			server.BladeDescr = getXMLAttributeData(body, "configResolveDn", "computeBlade", "descr")
			server.BladeDN = dn
			server.BladeLabel = getXMLAttributeData(body, "configResolveDn", "computeBlade", "usrLbl")
			server.BladeModel = getXMLAttributeData(body, "configResolveDn", "computeBlade", "model")
			server.BladeName = getXMLAttributeData(body, "configResolveDn", "computeBlade", "name")
			server.BladePID = getXMLAttributeData(body, "configResolveDn", "computeBlade", "partNumber")
			server.BladeSlot = getXMLAttributeData(body, "configResolveDn", "computeBlade", "slotId")
			server.BladeChassis = getXMLAttributeData(body, "configResolveDn", "computeBlade", "chassisId")
			server.BladeSerial = getXMLAttributeData(body, "configResolveDn", "computeBlade", "serial")
			server.BladeUUID = getXMLAttributeData(body, "configResolveDn", "computeBlade", "uuid")
			server.BladeMemory = getXMLAttributeData(body, "configResolveDn", "computeBlade", "totalMemory")
			server.BladeAssociation = getXMLAttributeData(body, "configResolveDn", "computeBlade", "association")
			server.BladeAssociatedTo = getXMLAttributeData(body, "configResolveDn", "computeBlade", "assignedToDn")
			server.BladeAvailability = getXMLAttributeData(body, "configResolveDn", "computeBlade", "availability")
			server.BladeSockets = getXMLAttributeData(body, "configResolveDn", "computeBlade", "numOfCpus")
			server.BladeCores = getXMLAttributeData(body, "configResolveDn", "computeBlade", "numOfCores")
			server.BladePower = getXMLAttributeData(body, "configResolveDn", "computeBlade", "operPower")
		}
	}
	u.LastResponse.Response = as.ToString(resp)
	u.LastResponse.Body = body
	u.LastResponse.Errors = err
	return server

}

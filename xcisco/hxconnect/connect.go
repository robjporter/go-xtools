package hxconnect

// GET /cluster/savings
// POST /clusterVersionDetails
// GET /virtplatform/cluster
// GET /virtplatform/vms
// GET /datastores
// GET /summary
// GET /appliances

import (
	"errors"
	"strings"
	"time"

	"../request"
)

type Connection struct {
	token       string
	timeout     int
	hxurls      *hxurls
	capurls     *capurls
	Metrics     *metrics
	Capital     *capital
	Credentials *Creds
}

type Creds struct {
	Url           string
	Username      string
	Password      string
	Client_id     string
	Client_secret string
}

type hxurls struct {
	Authentication    string
	About             string
	ClusterInfo       string
	ClusterSavings    string
	ClusterPlatform   string
	ClusterVM         string
	ClusterDatastores string
	ClusterSummary    string
	ClusterAppliances string
	ClusterVersion    string
}

type metrics struct {
	Server string
	Key    string
}

func init() {}
func New() *Connection {
	return &Connection{
		timeout: 30,
		hxurls:  getHXURLS(),
		capurls: getCapURLS(),
		Metrics: &metrics{},
		Capital: &capital{},
		Credentials: &Creds{
			Url:           "https://",
			Username:      "",
			Password:      "",
			Client_id:     "HxGuiClient",
			Client_secret: "Sunnyvale",
		},
	}
}

func getHXURLS() *hxurls {
	return &hxurls{
		Authentication:    "/aaa/v1/auth?grant_type=password", // POST
		About:             "/rest/about",                      // GET
		ClusterInfo:       "/rest/clusters",                   // GET
		ClusterSavings:    "/rest/cluster/savings",            // GET
		ClusterPlatform:   "/rest/virtplatform/cluster",       // GET
		ClusterVM:         "/rest/virtplatform/vms",           // GET
		ClusterDatastores: "/rest/datastores",                 // GET
		ClusterSummary:    "/rest/summary",                    // GET
		ClusterAppliances: "/rest/appliances",                 // GET
		ClusterVersion:    "/rest/clusterVersionDetails",      // POST
	}
}

func (c *Connection) SetToken(token string) {
	c.token = token
}

func (c *Connection) GetToken() string {
	return "Bearer " + strings.TrimSpace(c.token)
}

func (c *Connection) SetUsername(username string) {
	c.Credentials.Username = username
}

func (c *Connection) SetPassword(password string) {
	c.Credentials.Password = password
}

func (c *Connection) SetClientID(clientid string) {
	c.Credentials.Client_id = clientid
}

func (c *Connection) SetClientSecret(clientsecret string) {
	c.Credentials.Client_secret = clientsecret
}

func (c *Connection) SetUrl(url string) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		c.Credentials.Url = url
	} else {
		c.Credentials.Url = c.Credentials.Url + url
	}
}

func (c *Connection) SetTimeout(timeout int) {
	c.timeout = timeout
}

func (c *Connection) GetUsername() string {
	return c.Credentials.Username
}

func (c *Connection) GetPassword() string {
	return c.Credentials.Password
}

func (c *Connection) GetClientID() string {
	return c.Credentials.Client_id
}

func (c *Connection) GetClientSecret() string {
	return c.Credentials.Client_secret
}

func (c *Connection) GetUrl() string {
	return c.Credentials.Url
}

func (c *Connection) GetTimeout() time.Duration {
	return time.Duration(c.timeout) * time.Second
}

func (c *Connection) settingsMade() bool {
	if c.Credentials.Url != "https://" && c.Credentials.Username != "" && c.Credentials.Password != "" {
		return true
	}
	return false
}

func (c *Connection) sendPostRequestSimple(address string, url string, payload string) (*request.Client,error) {
	requester := request.New()
		_, err := requester.
			Post(address+url).
			Timeout(c.GetTimeout()).
			Send(payload).
			Set("content-type", "application/json").
			Set("cache-control", "no-cache").
			Accept("application/json").
			JSON()
		return requester,err
}

func (c *Connection) sendPostRequest(address string, url string, payload map[string]string) (*request.Client,error) {
	requester := request.New()
	_, err := requester.
		Post(address+url).
		Timeout(c.GetTimeout()).
		Send(payload).
		Set("content-type", "application/json").
		Set("cache-control", "no-cache").
		Accept("application/json").
		JSON()
	return requester,err
}

func (c *Connection) sendSecurePostRequest(address string, url string, payload map[string]string) (*request.Client,error) {
	requester := request.New()
	_, err := requester.
		Post(address+url).
		Timeout(c.GetTimeout()).
		Send(payload).
		Set("Authorization", c.GetToken()).
		Set("content-type", "application/json").
		Set("cache-control", "no-cache").
		Accept("application/json").
		JSON()
	return requester,err
}

func (c *Connection) sendGetRequest(address string, url string) (*request.Client,error) {
	requester := request.New()
	_, err := requester.
		Get(address+url).
		Timeout(c.GetTimeout()).
		Set("content-type", "application/json").
		Set("cache-control", "no-cache").
		Accept("application/json").
		JSON()
	return requester,err
}

func (c *Connection) sendSecureGetRequest(address string, url string) (*request.Client,error) {
	requester := request.New()
	_, err := requester.
		Get(address+url).
		Timeout(c.GetTimeout()).
		Set("Authorization", c.GetToken()).
		Set("content-type", "application/json").
		Set("cache-control", "no-cache").
		Accept("application/json").
		Text()
	return requester,err
}

func (c *Connection) GetResponseReason(req *request.Client) string {
	return req.ResponseReason()
}

func (c *Connection) GetResponseOK(req *request.Client) bool {
	return req.ResponseOK()
}

func (c *Connection) GetResponseURL(req *request.Client) string {
	return req.ResponseURL()
}

func (c *Connection) GetResponseCode(req *request.Client) int {
	return req.ResponseCode()
}

func (c *Connection) GetResponseData(req *request.Client) interface{} {
	return req.ResponseData()
}

func (c *Connection) GetResponseItem(req *request.Client,item string) interface{} {
	return req.ResponseDataItem(item)
}

func (c *Connection) GetResponseItemInt(req *request.Client,item string) int {
	return req.ResponseDataItemInt(item)
}

func (c *Connection) GetResponseItemInt64(req *request.Client,item string) int64 {
	return req.ResponseDataItemInt64(item)
}

func (c *Connection) GetResponseItemFloat(req *request.Client,item string) float64 {
	return req.ResponseDataItemFloat(item)
}

func (c *Connection) GetResponseItemString(req *request.Client,item string) string {
	return req.ResponseDataItemString(item)
}

func (c *Connection) GetResponseItemBool(req *request.Client,item string) bool {
	return req.ResponseDataItemBool(item)
}

func (c *Connection) GetResponseItemTime(req *request.Client,item string) time.Time {
	return req.ResponseDataItemTime(item)
}

func (c *Connection) Authenticate() (*request.Client,error) {
	if c.settingsMade() {
		data := map[string]string{"username": c.Credentials.Username, "password": c.Credentials.Password, "client_id": c.Credentials.Client_id, "client_secret": c.Credentials.Client_secret, "redirect_uri": c.Credentials.Url}
		req,e := c.sendPostRequest(c.Credentials.Url, c.hxurls.Authentication, data)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) About() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendGetRequest(c.Credentials.Url, c.hxurls.About)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterInfo() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterInfo)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterSavings() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterSavings)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterPlatform() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterPlatform)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterVM() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterVM)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterDatastores() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterDatastores)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterSummary() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterSummary)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterAppliances() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecureGetRequest(c.Credentials.Url, c.hxurls.ClusterAppliances)
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

func (c *Connection) ClusterVersion() (*request.Client,error) {
	if c.settingsMade() {
		req,e := c.sendSecurePostRequest(c.Credentials.Url, c.hxurls.ClusterVersion, make(map[string]string))
		return req,e
	}
	return nil,errors.New("Settings need to be updated before requests can be made.")
}

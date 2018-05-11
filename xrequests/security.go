package xrequests

import "crypto/tls"

func (a *Agent) TLSClientConfig(config *tls.Config) *Agent {
	a.Transport.TLSClientConfig = config
	return a
}

package xrequests

func (a *Agent) Get(url string) *Agent {
	a.ClearAgent()
	a.Method = GET
	a.URL = url
	return a
}

func (a *Agent) Post(url string) *Agent {
	a.ClearAgent()
	a.Method = POST
	a.URL = url
	return a
}

func (a *Agent) Head(url string) *Agent {
	a.ClearAgent()
	a.Method = HEAD
	a.URL = url
	return a
}

func (a *Agent) Put(url string) *Agent {
	a.ClearAgent()
	a.Method = PUT
	a.URL = url
	return a
}

func (a *Agent) Delete(url string) *Agent {
	a.ClearAgent()
	a.Method = DELETE
	a.URL = url
	return a
}

func (a *Agent) Patch(url string) *Agent {
	a.ClearAgent()
	a.Method = PATCH
	a.URL = url
	return a
}

func (a *Agent) Options(url string) *Agent {
	a.ClearAgent()
	a.Method = OPTIONS
	a.URL = url
	return a
}

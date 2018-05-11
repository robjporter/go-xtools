package main

import (
	"fmt"

	"../xipfilter"
)

func main() {
	ipfilters()
}

func ipfilters() {
	fmt.Println("")
	fmt.Println("SYSTEM IP FILTER **********************************************")
	options := xipfilter.Options{
		AllowedIPs:     []string{"10.52.208.1", "10.0.0.0/16"},
		BlockedIPs:     []string{},
		BlockByDefault: true,
	}
	filter := xipfilter.New(options)

	fmt.Println("IP FILTER - 10.0.0.1               >", filter.Allowed("10.0.0.1"))
	fmt.Println("IP FILTER - 10.0.42.1              >", filter.Allowed("10.0.42.1"))
	fmt.Println("IP FILTER - 10.42.0.1              >", filter.Allowed("10.42.0.1"))
	fmt.Println("IP FILTER - 10.52.208.1            >", filter.Allowed("10.52.208.1"))
	fmt.Println("IP FILTER - 10.52.208.10           >", filter.Allowed("10.52.208.10"))
}

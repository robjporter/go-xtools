package main

import (
	"../xhealth"
	"time"
	"fmt"
	"strings"
)

var (
	p *xhealth.Pinger
)

func main() {
	p = xhealth.New()
	p.NewURLMonitor("cisco", "http://www.cisco.com", 2)
	p.ChangedCallback("cisco", changed)

	p.NewURLMonitor("cisco2", "http://www.cisco.com", 3)
	p.ChangedCallback("cisco2", changed)

	p.NewURLMonitor("google", "http://www.google.com", 4)
	p.Callback("google", test)
	p.ChangedCallback("google", changed)

	p.NewPingMonitor("ad", "10.52.208.5", 3)
	p.Callback("ad", test2)

	fmt.Println("COUNT: ", p.Count())
	fmt.Println("ACTIVE COUNT: ", p.ActiveCount())
	p.StartMonitors()
	fmt.Println("ACTIVE COUNT: ", p.ActiveCount())
	time.Sleep(12 * time.Second)
	p.StopMonitors()
	fmt.Println("ACTIVE COUNT: ", p.ActiveCount())

	output()
}

func output() {
	names := p.GetMonitorNames()
	for i := 0; i < len(names); i++ {
		fmt.Println(strings.ToUpper(names[i]))
		fmt.Println("=======================")
		fmt.Println("Last Check: ", p.LastCheck(names[i]))
		fmt.Println("Last Status: ", p.LastStatus(names[i]))
		fmt.Println("Last Duration: ", p.LastDuration(names[i]))
		fmt.Println("Average Duration: ", p.AverageDuration(names[i]))
		fmt.Println("Total Successes: ", p.TotalSuccess(names[i]))
		fmt.Println("Total Failures: ", p.TotalFails(names[i]))
		fmt.Println("Total Checks: ", p.TotalChecks(names[i]))
		fmt.Println("")
	}
}

func test(name string, check int64, status bool, response string) {
	fmt.Println("=====================================================")
	fmt.Println("MONITOR: ", name)
	fmt.Println("LAST CHECK: ", check)
	fmt.Println("LAST STATUS: ", status)
	fmt.Println("=====================================================")
}

func test2(name string, check int64, status bool, response string) {
	fmt.Println("=====================================================")
	fmt.Println("MONITOR: ", name)
	fmt.Println("LAST CHECK: ", check)
	fmt.Println("LAST STATUS: ", status)
	fmt.Println("RESPONSE: ", response)
	fmt.Println("=====================================================")
}

func changed(name string, check int64, status bool, response string) {
	fmt.Println("=====================================================")
	fmt.Println(name, " has changed state to ", status)
	fmt.Println("=====================================================")
}

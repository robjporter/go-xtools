package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"../xsignals"
)

var s *xsignals.Signal

func main() {
	setupsignals()
	time.Sleep(20 * time.Second)
}

func setupsignals() {
	s = xsignals.New()
	s.Bind(syscall.SIGHUP, hupHandler)
	s.Bind(syscall.SIGTERM, termHandler)
	s.Bind(syscall.SIGINT, intHandler)
	s.Start()
}

func hupHandler() {
	fmt.Println("HUP")
	os.Exit(0)
	return
}

func termHandler() {
	fmt.Println("TERM")
	os.Exit(0)
	return
}

func intHandler() {
	fmt.Println("INT")
	os.Exit(0)
	return
}

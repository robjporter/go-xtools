package main

import (
	"fmt"

	"../xhttp"
	"time"
)

func main() {
	s := xhttp.NewSite()
	s.SetURL("http://www.google.co.uk")
	fmt.Println(s.Status())
	fmt.Println(s.LastStatus())
	fmt.Println(s.LastStatusString())

	// TODO:
	h := xhttp.NewHammer()
	h.SetThreadCount(4)
	h.SetHitCount(20)
	h.SetDelay(100*time.Millisecond)
	h.SetSpread(1*time.Second)
	h.SetURL("http://localhost:8080")
	h.Run()

	fmt.Println("HITS:     ",h.GetHits())
	fmt.Println("DURATION: ",h.GetDuration())
	fmt.Println("FAILED:   ",h.GetFailed())
	fmt.Println("QUICKEST: ",h.GetQuickest())
	fmt.Println("SLOWEST:  ",h.GetSlowest())
}

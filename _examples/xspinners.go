package main

import (
	"fmt"
	"time"

	"../xspinners"
)

func main() {
	spin := xspinners.New(xspinners.ARROWS1, " TEST")
	fmt.Println(spin)
	spin.Start()
	time.Sleep(10 * time.Second)
	spin.Stop()
}

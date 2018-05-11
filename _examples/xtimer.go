package main

import (
	"fmt"
	"time"

	"../xtimer"
)

func main() {

	xtimer.Timer("main")

	time.Sleep(5 * time.Second)

	fmt.Println(xtimer.Timer("main"))
}

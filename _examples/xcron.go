package main

import (
	"fmt"
	"time"

	"../xcron"
)

var cronjob, cronjob2 *xcron.CronJob

func main() {
	cronjob = xcron.New().SetName("Test").SetInterval(4).SetCallback(test)
	cronjob.Run()
	cronjob2 = xcron.New().SetName("Test2").SetInterval(6).SetCallback(test2)
	cronjob2.Run()

	time.Sleep(20 * time.Second)
}

func test() {
	fmt.Println("\n*******************************************************")
	fmt.Println("**           TIMED EVENT TRIGGERED NOW               **")
	fmt.Println("*******************************************************")
}

func test2() {
	fmt.Println("\n*******************************************************")
	fmt.Println("**           TIMED EVENT TRIGGERED NOW 2             **")
	fmt.Println("*******************************************************")
}

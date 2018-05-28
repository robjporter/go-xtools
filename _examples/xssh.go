package main

import (
	"fmt"
	"../xssh"
)

var (
	host     	= "IP"
	port     	= 22
	user       	= "USERNAME"
	password 	= "PASSWORD"
)

func main() {
	client, err := xssh.NewSShClient(host, int64(port), user, password)
	if err != nil {
		fmt.Println(err)
	}


	commands := []string{"show run | no-more","show run interface Ethernet107/1/2","show ver | no-more"}
	outPut,err := client.Commands(commands)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("COMMANDS: ",commands)
	fmt.Println("OUTPUT: ",outPut.String())

	client.Close()
}

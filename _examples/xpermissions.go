package main

import (
	"fmt"

	"../xpermissions"
)

func main() {
	anAcl := xpermissions.Acl{}
	anAcl.Grant("Admin", "Read", "Write", "Execute", "Water the flowers")
	fmt.Println(anAcl.Can("Admin", "Execute"))
	fmt.Println(anAcl.Can("Admin", "Water the flowers"))
	fmt.Println(anAcl.Can("Admin", "Test"))
	anAcl.Revoke("Admin", "Water the flowers")
	fmt.Println(anAcl.Can("Admin", "Water the flowers"))
}

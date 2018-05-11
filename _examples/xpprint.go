package main

import (
	"fmt"

	"../xpprint"
)

func main() {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println("Non pretty JSON                   >", x)
	xpprint.Pprint(x)
}

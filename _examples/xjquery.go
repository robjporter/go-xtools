package main

import (
	"fmt"

	"../xjquery"
)

var jsonData1 = `{"code":200,"message":"success","data":{ "busId":24, "mileage":253.56, "passenger":{ "students":[ [{"name":"Bili","age":16},{"name":"Celina","age":17},{"name":"Serafina","age":18}], [{"name":"Abby","age":19},{"name":"Amaris","age":20},{"name":"Fiona","age":21}], [{"name":"Snow","age":24},{"name":"Muse","age":23},{"name":"Gina","age":22}] ], "teachers":[{ "name":"Tom", "age":37, "teach":"math"}, { "name":"Li", "age":37, "teach":"math"} ] }} }`
var jsonData2 = `{"name":"ronin","age":24,"human":true,"object":{"x":1,"y":"bbb"},"array":[100,201,"TOP"]}`
var jsonData3 = `[{"type": "group","value": ["Lorem","Ipsum","dolor","sit",["A", "m", "e", "t"]]},{"type": "value","value": "Hello World"},{"type": "value","value": "foobar"}]`
var jsonData4 = `{"Name":"Eve","Age":6,"Parents":[{"Name":"Alice","Age":20},{"Name":"Bob","Age":21},{"Name":"Jane","Age":22},{"Name":"John","Age":23}]}`

func main() {
	example()
	//example1()
	//example2()
	//example3()
	//example4()
}

func example() {
	jq := gojsonq.New().JSONString(jsonData1)
	fmt.Println("STUDENTS 0:     ", jq.Find("data.passenger.students.[0]"))
	fmt.Println("STUDENTS 0 0:   ", jq.Reset().Find("data.passenger.students.[0].[0]"))
	fmt.Println("COUNT:          ", jq.Reset().From("data.passenger.students").Count())
	fmt.Println("COUNT:          ", jq.Reset().From("data.passenger.students.[0]").Count())
	fmt.Println("NAME 0.0:       ", jq.Reset().Find("data.passenger.students.[0].[0].name"))
	fmt.Println("AGES TOTAL:     ", jq.Reset().From("data.passenger.students.[0]").Sum("age"))
	fmt.Println("AGES AVERAGE:   ", jq.Reset().From("data.passenger.students.[0]").Avg("age"))
	fmt.Println("AGES MIN:       ", jq.Reset().From("data.passenger.students.[0]").Min("age"))
	fmt.Println("AGES MAX:       ", jq.Reset().From("data.passenger.students.[0]").Max("age"))
	fmt.Println("FIRST:          ", jq.Reset().From("data.passenger.students").First())
	fmt.Println("FIRST:          ", jq.Reset().From("data.passenger.students.[0]").First())
	fmt.Println("LAST:           ", jq.Reset().From("data.passenger.students").Last())
	fmt.Println("LAST:           ", jq.Reset().From("data.passenger.students.[0]").Last())
	fmt.Println("SORT:           ", jq.Reset().From("data.passenger.students.[0]").SortBy("age").Get())
	fmt.Println("Nth:            ", jq.Reset().From("data.passenger.students").Nth(1))
	fmt.Println("ONLY:           ", jq.Reset().From("data.passenger.students.[0]").Only("name").Get())
	fmt.Println("AGE < 18:       ", jq.Reset().From("data.passenger.students.[0]").Where("age", "<", 18).Get())
	fmt.Println("NAME = BILI:    ", jq.Reset().From("data.passenger.students.[0]").Where("name", "=", "Bili").Get())
	fmt.Println("WHERE CONTAINS I:", jq.Reset().From("data.passenger.students.[0]").WhereContains("name", "i").Get())
}

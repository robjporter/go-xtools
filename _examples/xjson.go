package main

import (
	"fmt"

	"../xjson"
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
	fmt.Println("== EXAMPLE =====================================================")
	x := xjson.New()
	x.ParseString(jsonData1)
	x.Get("data.passenger.students.0.0").ForEach(
		func(key string, value interface{}) bool {
			fmt.Println(key, value)
			return true
		})
	x.Get("data.passenger.students.0").ForEach(
		func(key string, value interface{}) bool {
			fmt.Println(key, value)
			return true
		})
	fmt.Println(x.Get("data.passenger.students.0.[age=16].name").StringArray())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(x.PrintTrim())
}

func example1() {
	fmt.Println("== EXAMPLE 1 =====================================================")
	x := xjson.New()
	x.ParseString(jsonData1)
	fmt.Println("LOADED: ", x.Exists())
	fmt.Println("CODE: ", x.Get("code").Int())
	fmt.Println("CODE EXISTS: ", x.Get("code").Exists())
	fmt.Println("RANDOM EXISTS: ", x.Get("random").Exists())
	fmt.Println("MESSAGE: ", x.Get("message").String())
	fmt.Println("MAP: ", x.Get("data").Map())
	fmt.Println("DATA COUNT: ", x.Get("data.#").Int())
	fmt.Println("DATA PASSENGER COUNT: ", x.Get("data.passenger.#").Int())
	fmt.Println("DATA BUDID: ", x.Get("data.busId").Int())
	fmt.Println("DATA MILEAGE: ", x.Get("data.mileage").Int())
	fmt.Println("DATA PASSENGER: ", x.Get("data.passenger").Interface())
	fmt.Println("DATA PASSENGER ISARRAY: ", x.Get("data.passenger.students.0.0.name").IsArray())
	fmt.Println("DATA PASSENGER STUDENTS ISARRAY: ", x.Get("data.passenger.students").IsArray())
	fmt.Println("DATA PASSENGER STUDENTS COUNT: ", x.Get("data.passenger.students.#").Int())
	fmt.Println("DATA PASSENGER TEACHERS COUNT: ", x.Get("data.passenger.teachers.#").Int())
	fmt.Println("DATA PASSENGER STUDENTS 0 COUNT: ", x.Get("data.passenger.students.0.#").Int())
	fmt.Println("DATA PASSENGER STUDENTS 0 0 NAME: ", x.Get("data.passenger.students.0.0.name").String())
	fmt.Println("DATA PASSENGER STUDENTS 0 0 NAME: ", x.Get("data.passenger.students.0.0.age").Int())
	fmt.Println("DATA PASSENGER STUDENTS 0 * NAME: ", x.Get("data.passenger.students.0.*.name").StringArray())
	fmt.Println("DATA PASSENGER STUDENTS 0 * AGE: ", x.Get("data.passenger.students.0.*.age").IntArray())
	fmt.Println("DATA PASSENGER STUDENTS 0 * AGE 1:2: ", x.Get("data.passenger.students.0.[1:2].age").IntArray())
	fmt.Println("DATA PASSENGER STUDENTS 0 * AGE :2: ", x.Get("data.passenger.students.0.[:2].age").IntArray())
	fmt.Println("DATA PASSENGER STUDENTS 0 * AGE 1:: ", x.Get("data.passenger.students.0.[1:].age").IntArray())
	fmt.Println("DATA PASSENGER STUDENTS 0 * NAME 1:: ", x.Get("data.passenger.students.0.[1:].name").StringArray())

	//fmt.Println(x.Print())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(x.PrintTrim())

}

func example2() {
	fmt.Println("== EXAMPLE 2 =====================================================")
	x := xjson.New()
	x.ParseString(jsonData2)
	fmt.Println("NAME: ", x.Get("name").String())
	fmt.Println("AGE: ", x.Get("age").Int())
	fmt.Println("HUMAN: ", x.Get("human").Bool())
	fmt.Println("OBJECT: ", x.Get("object").Interface())
	fmt.Println("OBJECT COUNT: ", x.Get("object.#").Int())
	fmt.Println("OBJECT 0 X: ", x.Get("object.x").Int())
	fmt.Println("OBJECT 0 Y: ", x.Get("object.y").String())
	fmt.Println("ARRAY: ", x.Get("array").Interface())
	fmt.Println("ARRAY COUNT: ", x.Get("array.#").Int())
	fmt.Println("ARRAY 0: ", x.Get("array.0").Int())
	fmt.Println("ARRAY 1: ", x.Get("array.1").Int())
	fmt.Println("ARRAY 2: ", x.Get("array.2").String())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(x.PrintTrim())
}

func example3() {
	fmt.Println("== EXAMPLE 3 =====================================================")
	x := xjson.New()
	x.ParseString(jsonData3)
	fmt.Println("COUNT ", x.Get("#").Int())
	fmt.Println("0: ", x.Get("0").Map())
	fmt.Println("0 TYPE: ", x.Get("0.type").String())
	fmt.Println("0 VALUE: ", x.Get("0.value").Interface())
	fmt.Println("0 VALUE 0: ", x.Get("0.value.0").String())
	fmt.Println("0 VALUE 0 ARRAY COUNT: ", x.Get("0.value.4.#").Int())
	fmt.Println("0 VALUE 0 ARRAY: ", x.Get("0.value.4").StringArray())
	fmt.Println("1 TYPE: ", x.Get("1.type").String())
	fmt.Println("2 TYPE: ", x.Get("2.type").String())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(x.PrintTrim())
}

func example4() {
	fmt.Println("== EXAMPLE 4 =====================================================")
	x := xjson.New()
	x.ParseString(jsonData4)
	fmt.Println("NAME: ", x.Get("Name").String())
	fmt.Println("AGE: ", x.Get("Age").Int())
	fmt.Println("PARENTS COUNT: ", x.Get("Parents.#").Int())
	fmt.Println("PARENTS: ", x.Get("Parents").Interface())
	fmt.Println("PARENTS 2:: ", x.Get("Parents.[2:]").Interface())
	fmt.Println("PARENTS 2: NAME: ", x.Get("Parents.[2:].Name").Interface())
	fmt.Println("PARENTS 0 NAME: ", x.Get("Parents.0").Interface())
	fmt.Println("PARENTS 0 NAME: ", x.Get("Parents.0.Name").String())
	fmt.Println("PARENTS 0 AGE: ", x.Get("Parents.0.Age").Int())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(x.PrintTrim())
}

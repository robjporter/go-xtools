package main

import (
	"../xregistry"
	"fmt"
	"time"
)

type info struct {
	name string
	age  int
}

func main() {
	i := &info{name: "test", age: 22}
	reg := xregsitry.New()

	fmt.Println(reg)

	reg.SetCustom("key1", 1, 1*time.Second)
	fmt.Println(reg.Get("key1"))
	time.Sleep(3 * time.Second)
	fmt.Println("HERE: ", reg.Get("key1"))
	reg.OnExpired = func(key interface{}, val interface{}) {
		fmt.Println("on expired", key, val)
	}
	reg.Close()

	reg.Set("key1", 1)
	reg.Set("key2", "test")
	reg.Set("key3", false)
	reg.Set("key4", 4.4)
	reg.Set("key5", i)
	reg.Set("key1", "1")
	reg.Set("key1.1", "1")
	reg.Set("key1.2", "2")
	reg.Set("key8.1", "1")
	reg.Set("key8.2", "2")

	fmt.Println("KEY1 :", reg.Get("key1"))
	fmt.Println("KEY1.1 :", reg.Get("key1.1"))
	fmt.Println("KEY1.2 :", reg.Get("key1.2"))
	fmt.Println("KEY2: ", reg.Get("key2"))
	fmt.Println("KEY3 :", reg.Get("key3"))
	fmt.Println("KEY4 :", reg.Get("key4"))
	fmt.Println("KEY5 :", reg.Get("key5"))
	fmt.Println("KEY6 :", reg.Get("key6"))
	reg.Close()
	/*
	i := &info{name: "test", age: 22}
	reg := xregistry.New()
	reg.Set("key1", 1)
	reg.Set("key2", "test")
	reg.Set("key3", false)
	reg.Set("key4", 4.4)
	reg.Set("key5", i)
	fmt.Println(reg)
	fmt.Println(reg.Get("key1"))
	reg.Set("key1", "1")
	reg.Set("key1.1", "1")
	reg.Set("key1.2", "2")
	reg.Set("key8.1", "1")
	reg.Set("key8.2", "2")
	fmt.Println(reg.Get("key1"))
	fmt.Println(reg.Get("key1.1"))
	fmt.Println(reg.Get("key1.2"))
	fmt.Println(reg.Get("key2"))
	fmt.Println(reg.Get("key3"))
	fmt.Println(reg.Get("key4"))
	fmt.Println(reg.Get("key5"))
	fmt.Println(reg.Get("key6"))
	fmt.Println(reg.Exists("key8"))
	fmt.Println(reg.Exists("key8.1"))
	fmt.Println(reg.Exists("key8.2"))
	reg.Copy("key8.1", "key8.3")
	fmt.Println(reg.Get("key8.3"))

	m := make(map[string]interface{})
	m["key10"] = "key10"
	m["key10.1"] = "key10.1"
	m["key10.1.1"] = "key10.1.1"
	m["key.10.1"] = "key.10.1"

	reg.SetWithMap(m)
	reg.Delete("key10")

	fmt.Println(reg.Keys())
	fmt.Println(reg.String())
	fmt.Println(reg.Len())
	*/
}

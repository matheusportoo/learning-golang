package main

import (
	"fmt"
	"reflect"
)

func main() {
	// make(map([<key-type>]<value-type>))
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3

	delete(m, "key3")

	fmt.Println(m, reflect.TypeOf(m))

	v1 := m["key1"]
	fmt.Println("v1: ", v1, reflect.TypeOf(v1))

	_, prs := m["key2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n["foo"])
}

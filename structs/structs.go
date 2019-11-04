package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 30})

	fmt.Println(newPerson("Matheus"))

	s := person{name: "Matheus", age: 30}

	sp := &s
	fmt.Println(sp.age)

	sp.age = 18
	fmt.Println(sp.age)
}

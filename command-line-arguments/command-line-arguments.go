package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println("Args with prog:", argsWithProg)
	fmt.Println("Args without prog:", argsWithoutProg)
	fmt.Println("Arg:", arg)
}

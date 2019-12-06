package main

import "os"

func main() {
	panic("a probleam")

	_, err := os.Create("./text.txt")
	if err != nil {
		panic(err)
	}
}

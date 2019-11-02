package main

import "fmt"

func intSeq(m int) func() int {
	i := 0

	return func() int {
		i++
		return i * m
	}
}

func main() {
	nextInt := intSeq(2)

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

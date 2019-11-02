package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	num := 7
	result := fact(num)
	fmt.Println("factorial of", num, "is", result)
}

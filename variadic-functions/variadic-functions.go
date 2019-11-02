package main

import "fmt"

func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("total:", total)
}

func main() {
	sum(1, 2, 3, 4, 5, 6)
	sum(10, 20, 30, 40, 50)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

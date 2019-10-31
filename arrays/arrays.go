package main

import "fmt"

func main() {
	var a [5]int
	a[4] = 100

	fmt.Println("a", a)

	//  syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("b", b)
	fmt.Println("size of b is:", len(b), "positions")

	fmt.Print("\n")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}

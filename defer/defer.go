package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("./defer.txt")
	s := "number"

	defer closeFile(f)

	for i := 0; i < 10; i++ {
		n := i
		r := fmt.Sprintf("%s: %d", s, n)
		writeFile(f, r)
	}

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File, s string) {
	fmt.Println("writing")
	fmt.Fprintln(f, s)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

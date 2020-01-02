package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 0, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	/*
		It’s also possible to declare an option that uses
		an existing var declared elsewhere in the program.
		Note that we need to pass in a pointer to the
		flag declaration function.
	*/

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	/*
		Once all flags are declared, call flag.Parse()
		to execute the command-line parsing.
	*/
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail", flag.Args())
}

/*
	go build -o bin/command-line-flags command-line-flags.go
	./bin/command-line-flags -word=opt -numb=7 -fork -svar=flag a b c d
*/

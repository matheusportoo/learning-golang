package main

import (
	"bufio"
	"fmt"
	"net/http"
	"flag"
)

func main() {
	urlPtr := flag.String("url", "https://gobyexample.com", "type -url=https://url.com")

	flag.Parse()

	resp, err := http.Get(*urlPtr)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

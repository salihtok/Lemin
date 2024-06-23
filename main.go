package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1]) != 2 {
		fmt.Println("Usage: go run . <mapFile>")
		os.Exit(1)
	}
	readFile(os.Args[1])

}

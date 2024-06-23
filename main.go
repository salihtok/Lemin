package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <mapFile>")
		os.Exit(1)
	}
	readFile(os.Args[1])
	SeperateRows() // verileri parçaladık.

}

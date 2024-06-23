package main

import (
	"fmt"
	"os"
)

func ContentError(startFlag, endFlag, connectFlag, commendFlag, antFlag bool) {
	if !startFlag {
		fmt.Println("No start room")
		os.Exit(1)
	}
	if !endFlag {
		fmt.Println("No end room")
	}
	if !connectFlag {
		fmt.Println("No connection")
		os.Exit(1)
	}
	if !commendFlag {
		fmt.Println("No comment room")
		os.Exit(1)
	}
	if !antFlag {
		fmt.Println("No ant count")
		os.Exit(1)
	}
}

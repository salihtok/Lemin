package main

import (
	"fmt"
	models "models/Models"

	"os"
)

func readFile(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	models.Content = string(data)
}

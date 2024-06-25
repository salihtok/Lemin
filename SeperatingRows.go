package main

import (
	models "models/Models"
	"strings"
)

func SeperateRows() {
	var temp string
	for _, row := range strings.Split(models.Content, "\n") {
		if row != "" {
			temp = strings.TrimSpace(row)
			models.Rows = append(models.Rows, strings.Split(temp, " "))
		}

	}
	//fmt.Println(models.Rows)
}

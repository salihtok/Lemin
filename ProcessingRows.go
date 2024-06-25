package main

import (
	"fmt"
	models "models/Models"
	"strconv"
	"strings"
)

func ProcessRows() {
	isThereStart := false // they can change in many way.
	isThereEnd := false
	isThereConnect := false
	isThereComment := false
	isThereAnts := false

	for i, row := range models.Rows {
		if len(row) == 1 && row[0] != "##start" && row[0] != "##end" {
			antCount, err := strconv.Atoi(row[0])
			if err == nil {
				models.AntsCount = antCount
				if models.AntsCount >= 1 {
					isThereAnts = true

				}
			}
		}
		if row[0] == "##start" {
			if len(models.Rows[i+1]) == 3 {
				models.StartRoom = models.Rows[i+1][0]
				isThereStart = true

			}
		}

		if row[0] == "##end" {
			if len(models.Rows[i+1]) == 3 {
				models.EndRoom = models.Rows[i+1][0]
				isThereEnd = true
			}
		}
		if len(row) == 3 && row[0] != models.StartRoom && row[0] != models.EndRoom {
			models.CommentRows = append(models.CommentRows, row[0])
			isThereComment = true
		}

		if len(row) == 1 && strings.Contains(row[0], "-") {
			models.ConnectRows = append(models.ConnectRows, row[0])
			isThereConnect = true
		}

	}
	ContentError(isThereStart, isThereEnd, isThereConnect, isThereComment, isThereAnts)
	FormatController()
	ConnectionController()

	fmt.Println("***************the test")
	fmt.Println(models.ConnectRows)
	fmt.Println("***************the test")
}

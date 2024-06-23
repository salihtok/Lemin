package main

import (
	"fmt"
	models "models/Models"
	"os"
	"strings"
)

func FormatController() {
	allRooms := append(models.CommentRows, models.StartRoom, models.EndRoom)

	for _, rows := range allRooms {
		if strings.HasPrefix(rows, "L") || strings.HasPrefix(rows, "#") {
			fmt.Println("Room name cannot start with L or #")
			os.Exit(1)
		}
	}
}

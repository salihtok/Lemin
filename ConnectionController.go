package main

import (
	"fmt"
	models "models/Models"
	"os"
	"strings"
)

func ConnectionController() {
	for _, connect := range models.ConnectRows {
		words := strings.Split(connect, "-")
		if !FindingRoom(words[0]) || !FindingRoom(words[1]) {
			fmt.Println("Connection error")
			os.Exit(1)
		}

	}
}
func FindingRoom(room string) bool {
	allRooms := append(models.CommentRows, models.StartRoom, models.EndRoom)
	for _, rows := range allRooms {
		if rows == room {
			return true
		}
	}
	return false
}

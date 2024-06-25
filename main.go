package main

import (
	"fmt"
	models "models/Models"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <mapFile>")
		os.Exit(1)
	}
	readFile(os.Args[1])
	SeperateRows() // verileri parçaladık.
	ProcessRows()
	fmt.Println("-------------------------------------------")
	fmt.Println("Odalar: ")
	fmt.Println(models.CommentRows)
	fmt.Println("Odaların bağlantıları: ")
	fmt.Println(models.ConnectRows)
	fmt.Println("Yollar :")
	fmt.Println(models.Roads)
	fmt.Println("ROWS as çektiğimiz bilgiler: ")
	fmt.Println(models.Rows)
	fmt.Println("-------------------------------------------")

	connections := models.ConnectRows
	graph := createGraph(connections)

	start := models.StartRoom
	end := models.EndRoom
	paths := graph.findAllPaths(start, end)

	if len(paths) == 0 {
		fmt.Println("Yol bulunamadı")
	} else {
		fmt.Println("Tüm olası yollar:")
		for _, path := range paths {
			fmt.Println(path)
		}
	}
}

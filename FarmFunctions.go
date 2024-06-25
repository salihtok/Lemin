package main

import (
	models "models/Models"
	"strings"
)

func FindingConnections(currentRoom string) []string {
	selectedRooms := []string{}
	for _, row := range models.ConnectRows {
		if strings.Contains(row, currentRoom+"-") || strings.Contains(row, "-"+currentRoom) {
			words := strings.Split(row, "-")
			for _, word := range words {
				if word != currentRoom {
					selectedRooms = append(selectedRooms, word)
				}
			}
		}
	}

	return selectedRooms

}

type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
	g.nodes[to] = append(g.nodes[to], from) // İki yönlü bağlantı
}

func createGraph(connections []string) *Graph {
	graph := NewGraph()
	for _, connection := range connections {
		parts := strings.Split(connection, "-")
		if len(parts) == 2 {
			graph.AddEdge(parts[0], parts[1])
		}
	}
	return graph
}
func (g *Graph) findAllPaths(source, sink string) [][]string {
	var paths [][]string
	var dfs func(current string, visited map[string]bool, path []string)

	dfs = func(current string, visited map[string]bool, path []string) {
		if visited[current] {
			return
		}
		visited[current] = true
		path = append(path, current)

		if current == sink {
			newPath := make([]string, len(path))
			copy(newPath, path)
			paths = append(paths, newPath)
		} else {
			for _, neighbor := range g.nodes[current] {
				if !visited[neighbor] {
					dfs(neighbor, visited, path)
				}
			}
		}

		visited[current] = false
		path = path[:len(path)-1]
	}

	dfs(source, make(map[string]bool), []string{})
	return paths
}

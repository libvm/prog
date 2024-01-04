package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type Graph struct {
	Adjacency map[string][]string `json:"adjacency"`
	Edges     map[string]Edge     `json:"edges"`
}

type Edge struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Weight int    `json:"weight"`
}

func readJSON(filename string) (Graph, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Graph{}, err
	}

	var graph Graph
	err = json.Unmarshal(data, &graph)
	if err != nil {
		return Graph{}, err
	}

	return graph, nil
}

func dijkstra(graph Graph, start, end string) (int, []string) {
	dist := make(map[string]int)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	for node := range graph.Adjacency {
		dist[node] = math.MaxInt64
		prev[node] = ""
	}
	dist[start] = 0

	for {
		current := ""
		minDist := math.MaxInt64
		for node, d := range dist {
			if !visited[node] && d < minDist {
				current = node
				minDist = d
			}
		}

		if current == "" || dist[current] == math.MaxInt64 {
			break
		}

		visited[current] = true

		for _, neighbor := range graph.Adjacency[current] {
			edge := graph.Edges[neighbor]
			alt := dist[current] + edge.Weight
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = current
			}
		}
	}

	path := []string{end}
	current := end
	for prev[current] != "" {
		current = prev[current]
		path = append([]string{current}, path...)
	}

	return dist[end], path
}

func main() {
	filename := "graph.json"
	graph, err := readJSON(filename)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	startNode := "a"
	endNode := "e"
	distance, path := dijkstra(graph, startNode, endNode)

	if distance < math.MaxInt64 {
		fmt.Printf("Shortest path from %s to %s: %v\n", startNode, endNode, path)
		fmt.Printf("Total distance: %d\n", distance)
	} else {
		fmt.Printf("No path found from %s to %s\n", startNode, endNode)
	}
}

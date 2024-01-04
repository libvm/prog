package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Options []Option `json:"options"`
}

type Option struct {
	Text   string `json:"text"`
	NextId int    `json:"next"`
}

func readJSON(filename string) ([]Node, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var nodes []Node
	err = json.Unmarshal(content, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func playQuest(nodes []Node) {
	currentNode := nodes[0]

	for {
		fmt.Println(currentNode.Text)

		if len(currentNode.Options) == 0 {
			fmt.Println("Конец игры.")
			break
		}

		fmt.Println("Ваши варианты:")
		for i, option := range currentNode.Options {
			fmt.Printf("%d. %s\n", i+1, option.Text)
		}

		var choice int
		fmt.Print("Выберите вариант: ")
		fmt.Scan(&choice)

		if choice < 1 || choice > len(currentNode.Options) {
			fmt.Println("Некорректный выбор. Пожалуйста, выберите снова.")
			continue
		}
		currentNode = nodes[currentNode.Options[choice-1].NextId-1]
	}
}
func main() {
	nodes, err := readJSON("game.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	playQuest(nodes)
}

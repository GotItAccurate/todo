package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const TodoFile = "todos.json"

type Todo struct {
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"createdAt"`
	Deadline  time.Time `json:"deadline"`
	Done      bool      `json:"done"`
}

func loadTodos() []Todo {
	file, err := os.ReadFile(TodoFile)
	if err != nil {
		return []Todo{}
	}

	var todos []Todo
	json.Unmarshal(file, &todos)

	return todos
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", " ")
	os.WriteFile(TodoFile, data, 0644)
}

func PrintTodos(todos []Todo) {
	for i, v := range todos {
		fmt.Printf("%d. %s - %v\n", i+1, v.Task, v.Done)
	}
}

package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text   string    `json:"content"`
}

func (todo Todo) Display() {
	fmt.Printf("The todo is: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err1 := json.Marshal(todo)
	if err1 != nil {
		return err1
	}

	err2 := os.WriteFile(fileName, json, 0644)
	return err2
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("empty string passed in as input")
	}
	return Todo{content}, nil
}

package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"` // struct field tags for JSON serialization - metadata
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid Input!")
	}
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("\nTodo: '%v'\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonData, 0644)

	if err != nil {
		return err
	}

	fmt.Println("Todo saved successfully!")
	return nil
}

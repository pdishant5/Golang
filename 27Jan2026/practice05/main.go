package main

import (
	"bufio"
	"fmt"
	"noteApp/note"
	"os"
	"strings"
	"todoApp/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// defining the interfaces separately
// type outputtable interface {
// 	save() error
// 	Display()
// }

// alternatively, embedding interfaces to create a new interface
type outputtable interface {
	saver
	// displayer // or
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todoItem, err := todo.New(todoText)

	if err != nil {
		fmt.Printf("Error creating todo: %v\n", err)
		return
	}

	err = DisplayData(todoItem)

	if err != nil {
		return
	}

	noteItem, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = DisplayData(noteItem)

	if err != nil {
		return
	}
}

func PrintSomething(value interface{}) { // empty interface can accept value of any type
	// type switch
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	default:
		fmt.Println("Unknown type")
	}

	// alternate way using type assertion
	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer using type assertion:", intValue)
		return
	}

	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("Float using type assertion:", floatValue)
		return
	}

	strValue, ok := value.(string)
	if ok {
		fmt.Println("String using type assertion:", strValue)
		return
	}
}

func DisplayData(data outputtable) error {
	data.Display()
	return SaveData(data)
}

func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return err
	}

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Remove the newline character from the input
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

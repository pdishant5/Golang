package main

import (
	"bufio"
	"fmt"
	"noteApp/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	noteItem, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	noteItem.Display()
	err = noteItem.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
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

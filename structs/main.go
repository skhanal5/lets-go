package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/note"
)

func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if (err != nil) {
		return
	}
	note.Display()
	err = note.Save()
	if (err != nil) {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving was successful")
}

func getUserInput(prompt string) string{
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if (err != nil) {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows
	return text
}
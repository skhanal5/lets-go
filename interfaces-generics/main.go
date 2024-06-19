package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"interfaces/note"
	"interfaces/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	Display()
	saver		//embedding existing interface into another interface
}

func main() {
	printSomething("a")
	printSomething(1)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if (err != nil) {
		return
	}

	note, err := note.New(title, content)
	if (err != nil) {
		return
	}

	outputData(todo)
	err = saveData(todo)

	if err != nil {
		return 
	}

	outputData(note)
}

func printSomething(value interface{}) {	// any value is allowed when you specify the type as interface{}
	switch value.(type)	{ 	//switching on the type of the variable
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println(value)
	default:
		// returns typedVal (value in that type), a boolean indicating if it is that type
		typedValue, ok := value.(int)
		if ok {
			typedValue += 1
			fmt.Println("Integer:", typedValue)
		}
		fmt.Println(typedValue)	// will be the null value if the check fails, so 0 
	}	
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData (data saver) error {
	err := data.Save()
	if (err != nil) {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving was successful")
	return nil
}


func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
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
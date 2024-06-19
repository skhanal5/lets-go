package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`		//struct tags = metadata that can be used by other packages
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("The title is: %s\nThe content is: %s\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"

	json, err1 := json.Marshal(note)
	if err1 != nil {
		return err1
	}

	err2 := os.WriteFile(fileName, json, 0644)
	return err2
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("empty string passed in as input")
	}
	return Note{title, content, time.Now()}, nil
}

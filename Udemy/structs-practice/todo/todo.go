package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func NewNote(text string) *Todo {
	return &Todo{
		Text: text,
	}
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

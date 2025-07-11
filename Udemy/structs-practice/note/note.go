package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewNote(title, description string) *Note {
	return &Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func (note Note) Display() {
	fmt.Printf("Title Is %v, Description Is %v, Created At %v", note.Title, note.Description, note.CreatedAt.Format(time.RFC3339))
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")

	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

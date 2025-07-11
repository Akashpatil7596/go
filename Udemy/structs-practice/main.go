package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func getNoteData() (string, string, error) {
	noteTitle, err := getUserInput("Enter a note title: ")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	noteDescription, err := getUserInput("Enter a note description: ")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return noteTitle, noteDescription, nil
}

func getToDoData() string {
	text, _ := getUserInput("Enter Text:")
	return text

}

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

func printSomething(value interface{}) {

	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer value:", typedVal+1)

	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println("String value:", value)
	// case float64:
	// 	fmt.Println("Float value:", value)
	// case int:
	// 	fmt.Println("Integer value:", value)
	// }

}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)

}

func main() {
	printSomething(1.5)
	printSomething("Hello, World!")
	title, description, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	text := getToDoData()

	todo := todo.NewNote(text)

	err = outputData(*todo)

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}

	newNote := note.NewNote(title, description)

	err = outputData(*newNote)

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}

	fmt.Println("Note saved successfully!")

}

func getUserInput(prompt string) (string, error) {

	fmt.Printf("%v ", prompt)

	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	if text == "" {
		return "", errors.New("empty")
	}

	return text, nil
}

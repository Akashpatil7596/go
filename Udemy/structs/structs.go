package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var userDetail *user.User

	userDetail, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("ap3135198@gmail.com", "12345")
	admin.GetAdminDetailOutput()

	// var userDetailTwo user = user{
	// 	firstName: "AS",
	// 	lastName:  "PA",
	// 	birthdate: "0000",
	// 	createdAt: time.Now(),
	// }

	userDetail.GetUserDetailOutput()
	userDetail.ClearUserName()

	userDetail.GetUserDetailOutput()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

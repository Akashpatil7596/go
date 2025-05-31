package main

import (
	"errors"
	"fmt"

	"example.com/bank/fileOperations"
	"github.com/Pallinder/go-randomdata"
)

func main() {

	fmt.Println(randomdata.SillyName())

	var arr = [...]int{1, 2, 3, 4, 5, 2, 1, 4}
	fmt.Println("arr", arr)

	var slicedArray = arr[4:7]
	fmt.Println("slicedArray", slicedArray)

	var slice = []int{1, 4, 5, 7, 3}
	slice = append(slice, 1, 12)
	fmt.Println("slice", slice)

	subjectMarks := map[string]int{"Maths": 12, "English": 76, "Science": 54}

	rollNumbers := map[int]string{1: "Alex", 2: "John Doe"}

	rollNumbers[4] = "Allie"

	delete(rollNumbers, 2)

	fmt.Println(subjectMarks["Maths"])

	fmt.Println(rollNumbers)

	for number, squared := range rollNumbers {
		fmt.Printf("Square of %d is %s\n", number, squared)
	}

	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Programiz" {
		fmt.Println(myError)
	}

	accountBalance, err := fileOperations.GetFloatFromFile("balance.txt")

	if err != nil {
		fmt.Println(err)
	}

	for {

		var depositedMoney float64
		var withdrawnMoney float64

		displayWelcomeMessage()
		Hello()
		var userChoice int

		fmt.Println("What do you want to do?")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2:
			fmt.Print("Deposit Money : ")
			fmt.Scan(&depositedMoney)

			if depositedMoney <= 0 {
				fmt.Println("Ye nai chalega vro")
				continue
			}

			accountBalance += depositedMoney
			fmt.Println("Total Balance: ", accountBalance)
			fileOperations.WriteFloatToFile(accountBalance, "balance.txt")
		case 3:
			fmt.Print("How much money do you want? : ")
			fmt.Scan(&withdrawnMoney)

			if withdrawnMoney <= 0 {
				fmt.Println("Ye nai chalega vro")
				continue
			}

			if withdrawnMoney > accountBalance {
				fmt.Println("Ye nai chalega vro")
				continue
			}

			accountBalance -= withdrawnMoney
			fmt.Println("Total Balance: ", accountBalance)

			fileOperations.WriteFloatToFile(accountBalance, "balance.txt")
		default:
			fmt.Println("Bye Bye!")
			fmt.Println("Bye Love")
			// return
			return
		}

		// if userChoice == 1 {
		// 	fmt.Println(accountBalance)
		// } else if userChoice == 2 {
		// 	fmt.Print("Deposit Money : ")
		// 	fmt.Scan(&depositedMoney)

		// 	if depositedMoney <= 0 {
		// 		fmt.Println("Ye nai chalega vro")
		// 		continue
		// 	}

		// 	accountBalance += depositedMoney
		// 	fmt.Println("Total Balance: ", accountBalance)
		// } else if userChoice == 3 {
		// 	fmt.Print("How much money do you want? : ")
		// 	fmt.Scan(&withdrawnMoney)

		// 	if withdrawnMoney <= 0 {
		// 		fmt.Println("Ye nai chalega vro")
		// 		return
		// 	}

		// 	if withdrawnMoney > accountBalance {
		// 		fmt.Println("Ye nai chalega vro")
		// 		return
		// 	}

		// 	accountBalance -= withdrawnMoney
		// 	fmt.Println("Total Balance: ", accountBalance)
		// } else {
		// 	fmt.Println("Bye Bye!")
		// 	// return
		// 	break
		// }
	}

}

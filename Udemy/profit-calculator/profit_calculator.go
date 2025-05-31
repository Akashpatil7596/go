package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	profit := (revenue - expenses) * (1 - taxRate/100)

	fmt.Printf("Profit: %.2f\n", profit)
}

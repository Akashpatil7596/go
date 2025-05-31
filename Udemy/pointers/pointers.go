package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	// dereferencing the pointer
	fmt.Println(agePointer)

	// fmt.Println(age)

	adultYears := getAdultsYear(agePointer)

	fmt.Println(adultYears)
	fmt.Println(age)

}

func getAdultsYear(age *int) int {
	// return *age - 18

	*age = *age - 18

	return *age
}

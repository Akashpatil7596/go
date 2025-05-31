package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	var investmentAmount float64

	var years float64

	expectedReturnRate := 5.5

	outputText("Type your Investment amount: ")

	fmt.Scan(&investmentAmount)

	outputText("Type the number of years: ")

	fmt.Scan(&years)

	outputText("Type the expected return rate: ")

	fmt.Scan(&expectedReturnRate)

	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value is %.2f\n", futureValue)

	fmt.Println("formattedFV", formattedFV)

	formattedRFV := fmt.Sprintf("Future real value is %.2f\n", futureRealValue)

	fmt.Println("formattedRFV", formattedRFV)

	// fmt.Printf(`Future value is: %.2f\n
	// Future real value is %.2f\n`, futureValue, futureRealValue)

	// fmt.Println("Future value is", futureValue)

	// fmt.Printf("Future value is %T\n", futureValue)

	// fmt.Printf("Future value is %.2f\n", futureValue)

	// fmt.Println("Future real value is", futureRealValue)

	_, error := sum(1, 2)

	if error != nil {
		fmt.Println("error", error)
	}

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, years, expectedReturnRate float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
}

func sum(a, b int) (int, error) {
	res := a + b
	return fmt.Print(res)
}

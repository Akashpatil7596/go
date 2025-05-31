package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5 

func main() {
	
	var investmentAmount float64

	var years float64

	var expectedReturnRate float64

	fmt.Print("Investment Anount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate Years: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)	

	// * fmt.Println("Future Real Value Is : ",futureRealValue)

	formattedFV := fmt.Sprintf("Future Value Is %.2f", futureValue)

	formattedFRV := fmt.Sprintf("Future Real Value Is %.2f", futureRealValue)

	outputText(formattedFV, formattedFRV)

}

func outputText(output1 string,output2 string){
	fmt.Print(output1 , "\n", output2, "\n")
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue =  futureValue / math.Pow(1 + inflationRate / 100, years)
	
	return futureValue,futureRealValue
}
package fileOperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
		// panic(err.Error())
	}

	balanceText := string(data)

	balance, _ := strconv.ParseFloat(balanceText, 64)

	// if err != nil {
	// 	panic(err)
	// }

	return balance, nil
}

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

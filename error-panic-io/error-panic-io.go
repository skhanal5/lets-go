package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Input revenue: ")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserInput("Input expenses: ")
	if err != nil {
		panic(err)
	}
	taxRate, err := getUserInput("Input tax rate: ")
	if err != nil {
		panic(err)
	}

	earningsBeforeTax, earningsAfterTax, ratio := calc(revenue, expenses, taxRate)

	outputStr := fmt.Sprintf("EBT %f \nProfit %f \nRatio %f \n", earningsBeforeTax, earningsAfterTax, ratio)

	os.WriteFile("output.txt", []byte(outputStr), 0644)
	fmt.Print(outputStr)

}

func getUserInput(input string) (float64, error) {
	value := 0.0
	fmt.Print(input)
	fmt.Scan(&value)
	if value < 0 {
		return 0, errors.New("Non-negative values not allowed")
	} else if value == 0 {
		return 0, errors.New("Zero value not allowed")
	}
	return value, nil
}

func calc(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsAfterTax := (revenue - expenses) * (1 - taxRate/100)
	earningsBeforeTax := (revenue - expenses)
	ratio := earningsAfterTax / earningsBeforeTax
	return earningsBeforeTax, earningsAfterTax, ratio
}

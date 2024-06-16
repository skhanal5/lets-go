package main

import "fmt"

func main() {
	revenue := getUserInput("Input revenue: ")
	expenses := getUserInput("Input expenses: ")
	taxRate := getUserInput("Input tax rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calc(revenue,expenses,taxRate)

	fmt.Printf("EBT %f \n", earningsBeforeTax)
	fmt.Printf("Profit %f \n", earningsAfterTax)
	fmt.Printf("Ratio %f \n", ratio)
	
}

func getUserInput(input string) float64{
	revenue := 0.0
	fmt.Print(input)
	fmt.Scan(&revenue)
	return revenue
}


func calc(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsAfterTax := (revenue-expenses) * (1-taxRate/100)
	earningsBeforeTax := (revenue-expenses)
	ratio := earningsAfterTax/earningsBeforeTax
	return earningsBeforeTax, earningsAfterTax, ratio 
}
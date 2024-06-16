package main

import "fmt"

func main() {
	revenue, expenses, tax_rate := 0.0, 0.0, 0.0

	fmt.Print("Input revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input tax_rate: ")
	fmt.Scan(&tax_rate)
	
	earnings_after_tax := (revenue-expenses) * (1-tax_rate/100)
	earnings_before_tax := (revenue-expenses)

	fmt.Printf("EBT %f \n", earnings_before_tax)
	fmt.Printf("Profit %f \n", earnings_after_tax)
	fmt.Printf("Ratio %f \n", (earnings_before_tax/(earnings_after_tax)))
	

}
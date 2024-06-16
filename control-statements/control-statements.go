package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			
			if (depositAmount < 0) {
				fmt.Printf("Invalid amount: %f \n", depositAmount)
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if (withdrawAmount > accountBalance || withdrawAmount < 0) {
				fmt.Printf("Invalid amount: %f \n", withdrawAmount)
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Pce")
			break;
		}
	}
	fmt.Println("Thanks for using this bank.")
}
package main

import (
	"errors"
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan bool, len(taxRates))


	for index, taxRate := range taxRates {
		
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan errors)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i, _ := range taxRates {
		/*
			Waits for channel that emits the value first, ignores all other channels
		*/
		select {
		case err := <- errorChans[i]:
			fmt.Println(err)
		case <- doneChans[i]:
		}  
	}

	for _, doneChan := range doneChans {
		<- doneChan
	}
}

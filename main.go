package main

import (
	"fmt"
)

func main() {
	var rev float64
	var exp float64
	var tr float64

	fmt.Print("Revenue: ")
	fmt.Scan(&rev)

	fmt.Print("Expenses: ")
	fmt.Scan(&exp)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&tr)

	ebt := rev - exp
	tax := ebt * (tr / 100)
	prof := ebt - tax
	rat := (prof / rev) * 100

	fmt.Printf("Ebt: $%.2f", ebt)
	fmt.Printf("\nTax bill was: $%.2f", tax)
	fmt.Printf("\nProfit after tax: $%.2f", prof)
	fmt.Printf("\nProfit to revenue rate: %%%.2f", rat)
}

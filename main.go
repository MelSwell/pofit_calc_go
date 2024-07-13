package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	rev, _ := strconv.Atoi(os.Args[1])
	exp, _ := strconv.Atoi(os.Args[2])
	tr, _ := strconv.ParseFloat(os.Args[3], 64)

	e := float64(rev - exp)
	p := e - (e * tr)
	r := e / p

	fmt.Printf("Ebt: $%.2f", e)
	fmt.Printf("\nProphet after tax: $%.2f", p)
	fmt.Printf("\nEarnings to prophet rate: %.2f", r)
}

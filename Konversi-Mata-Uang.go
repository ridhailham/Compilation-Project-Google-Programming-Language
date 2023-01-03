package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	number, err := strconv.ParseFloat(arg[0], 64)

	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	eur := 0.88
	gbp := 0.78
	jpy := 113.02

	resultEur := number * eur
	resultGbp := number * gbp
	resultJpy := number * jpy

	fmt.Printf("%.2f USD is %.2f EUR\n", number, resultEur)
	fmt.Printf("%.2f USD is %.2f GBP\n", number, resultGbp)
	fmt.Printf("%.2f USD is %.2f JPY\n", number, resultJpy)

}

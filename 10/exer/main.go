package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollors")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is:", tax, "dollors")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is:", total, "dollors.")
	var availableFunds int = 120
	fmt.Println("Available funds", availableFunds, "dollors available")
	fmt.Println("within budget?", total <= float64(availableFunds))
}

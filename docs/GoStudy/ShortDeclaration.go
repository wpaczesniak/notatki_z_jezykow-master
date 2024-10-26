package main

import (
	"fmt"
)

func main() {
	// quantity := 10
	// length, width := 10.5, 5.2
	// customerName := "Damon Cole"

	// // fmt.Println("Quantity is", quantity)
	// fmt.Println("Size is", length, "by", width)
	// fmt.Println("Customer is", customerName)
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Powierczhnia wynosi", length*float64(width))
	fmt.Println("length > width?", length > float64(width))

}

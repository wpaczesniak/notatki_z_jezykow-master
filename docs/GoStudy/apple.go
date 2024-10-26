package main

import (
	"fmt"
)

func main() {
	var orginalCount int
	var eatenCount int

	orginalCount = 10
	eatenCount = 4

	fmt.Println("I started with", orginalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", orginalCount-eatenCount, "apples left.")
}

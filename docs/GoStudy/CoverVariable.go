package main

import (
	"fmt"
)

func main() {
	var count int = 12
	var suffix string = "minut dadatkowych materiałów"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol", "Francais")
	fmt.Println(count, suffix, format, languages)
}

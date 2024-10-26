package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# t# super język"
	replacer := strings.NewReplacer("#", "o")
	var fixed string = replacer.Replace(broken)
	fmt.Println(fixed)
}

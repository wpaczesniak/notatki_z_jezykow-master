package main

import "fmt"

// Zmienna globalna, dostępna we wszystkich funkcjach w pakiecie main
var globalVar = "To jest zmienna globalna"

func main() {
	// Zmienna lokalna dla funkcji main
	var functionVar = "To jest zmienna funkcji main"

	fmt.Println(globalVar)   // Dostęp do zmiennej globalnej
	fmt.Println(functionVar) // Dostęp do zmiennej funkcji main

	// Zasięg blokowy - zmienna dostępna tylko wewnątrz tego bloku if
	if true {
		var blockVar = "To jest zmienna bloku if"
		fmt.Println(blockVar) // Zmienna blockVar jest dostępna w if
	}
	fmt.Println(blockVar) // Dostęp do zmiennej globalnej
	// Próba dostępu do blockVar tutaj skończy się błędem, bo jest poza zakresem
	// fmt.Println(blockVar) // Błąd: blockVar nie jest zdefiniowana
}

// 	anotherFunction()
// }

// func anotherFunction() {
// 	// Zmienna lokalna dla funkcji anotherFunction
// 	var functionVar = "To jest zmienna funkcji anotherFunction"

// 	fmt.Println(globalVar)   // Dostęp do zmiennej globalnej
// 	fmt.Println(functionVar) // Dostęp do zmiennej lokalnej anotherFunction
// 	// Próba dostępu do zmiennej functionVar z main lub blockVar spowodowałaby błąd
// }

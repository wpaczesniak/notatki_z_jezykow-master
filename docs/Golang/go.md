# Notatki języka Go language


## Źródła: 

Dany opis został napisany w oparciu o książkę *Rusz Głową - Go*.

Język Go powstał w 2009 roku w oparciu o licencję Open source przez Googla. Prace nad językiem zaczely się 2007 roku. Głownym zadaniem bylo eliminacja długiego oczekiwania na zakonczenie kompilacji. W założeniach nowy projekt stworzenia języka opierał się na poniższych celach:

- szybka kompilacja,
- mniej skomplikowany kod,
- automatyczne zwalnianie nieużywanej pamięci (umożliwienie ponownego używania),
- łatwe pisanie oprogramowania wykonujące różne operacje(współbieżność),
- dobra obsługa procesów wielordzeniowych.

## Pierwszy program w go


```go
package main

import "fmt"

func main(){
    fmt.Println("Hello, Go!")
}
```

Kod w języku Go dzieli się na trzy części:

1) Klauzura *package*
2) instrukcje *import*
3) wykonywany kod
   
Pakiet to zbiór kodu oprzeznaczonego do podobnych zadań.
Pakiet **main** służy do tego, aby określić, że dany plik jest programem wykonywalnym, a nie tylko biblioteką. Pakiet **main** mówi że kod ma być uruchamiany bezpośrednio.

W każdym pliku należy zaimportować inne pakiety, aby kod z danego pliku mógł używać kodu z podanych pakietów. Nalezy unikać wczytywania całego kodu języka **Go** co może powodować powstanie dużego i powolnego programu. Należy importować tylko te pakiety które będą używane. 

Ostatnia część to wykonywany kod: 
Ta część podzielona jest na funkcje. Funkcja to grupa wierszy, które można wykonywać w innych miejscach programu. 

Dodatkowe informacje:

W języku Go dopuszczalne jest używanie średników ale nie jest to wymagane ale nie jest dobrze widziane. 

Jezyk go dostarcza narzedzia ktore pozwalają na automatyczne formatowanie kodu. Wykorzystuje się do tego komendy *go format*.

## Komunikaty o błędach

dokonując drobnej modyfikacji pierwszego kodu tak by otrzymać komunikat o błedzie. Pomijamy nawiasy.

```go
package main

import "fmt"

func main(){
    fmt.Println "Hello, Go!"
}
```

./test.go:6:13: syntax error: unexpected literal "Hello, World!" at end of statement

Język Go informuje w łatwy sposób informuje o błędzie. 

## Funkcja *Println*

Funkcję *Println* może przyjmować dowolną liczbę argumentów nawet zero. Wyświetlane argumenty są wyświetlane w oknie terminala z rozdzieloną spacją. Człon *ln* oznacza przejście do kolejnego wiersza w oknie terminala. 

Aby uzyskać dostęp do danej funkcji należy zaimportować zawierający ją pakiet. 

Dla przykładu 

**fmt.Println** 
- *fmt* → nazwa pakietu
- *Println()* → nazwa funkcji


Aby zaimportować funkcje z kilku różnych pakietów należy umieszcic w nawiasie. Należy pamiętać że jeden pakiet w jednej linii.

```go
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("hello"))
}
```

## Typy w Go

- Literał tekstowy *(ang. string literal)* - ciąg znaków; tekst pomiedzy cudzysłowami.
- Runy *(ang. rune)* - oznaczaja pojedyńcze znaki. Przechowywane sa jako kody liczbowe(standard Unicode) Przekazując runę do funkcji *fmt.Println* otrzymamy kod a nie znak. 
- wartości logiczne - true/false 
- liczby - Język Go traktuje liczby calkowite i zmiennoprzecinkowe jako wartości roznych typów.
- Operacje matematyczne tak jak to jest we wszystkich innych językach:
    - == sluży do sprawdzenia czy wartosci są równe,
    - != służy do sprawdzenia czy wartosci są rożne,
    - < = > wieksza, równa
  
## Statyczna kontrola typów

Satyczna analiza typów oznacza że jeszcze przed uruchomieniem programu wie, jakiego typu będą wartości. Funkcje oczekują argumentow określonych typów, a warości zwracane mają ustalony typ. Język Go jeszcze przed uruchomieniem wyswietli komunikat o błedzie, pozwala to uniknąć problemu, zanim zauważą go użytkownicy. 

## Deklarowanie zmiennych

Zmienna to fragment pamięci zawierający wartość. 
W języku Go aby poprawnie zadeklarowac zmienna nalezy najpierw podać słowo kluczowe **var**, nastepnie **nazwę zmiennej** i **typ** przechowywanej wartości.
Można deklarować jednym wierszu wiele zmiennych tego samego typu.

```go
var quality int
var length, width float64
var customerName string
```
Pożniej można przejśc do przypisania wartości uzywajac operatora **=**. 

```go
quantity = 2
customerName = "Joanna"
```

W jednej linii można przypisaćwiele wartosci do wielu zmiennych. Nazwy te nalezy oddzielić je między sobą przecinkami. Liczba zmiennych musi pokrywac się z liczbą wartości które chcemy do nich przypisać. Jeśli nie będzie równa to wyświetli się błąd. 

```go
length, width = 1.2, 2.4
```

Można oddrazu deklarować zmienne i przypisac im wartosci.

```go
var quantily int = 4
```

Można przypisac do zmiennej nowe wartości, jednak muszą być to wartości tego samego typu. Wynika to wcześniej wspomnianej statycznej kontroli typów. 

Jeśli zadeklaurujemy i nie przypiszemy żadnych wartosci to zmiennych to te zmienne będą miały przypisaną wartość zerową dla danego typu.

- 0 dla zmiennych typu int, float64 itd.
- false dla zmiennych typu bool 
- pusty łancuch znaków dla string.

### Krótkie deklarowanie zmiennych

Jeśli jednak zaraz po zadeklarowaniu zmiennej znana jest jej poczatkowa wartość, częściej stosuje się **krótkie deklaracje zmiennych**. Zamiast jawnego deklarowania typu zmiennej oraz przypisywania do nich wartości za pomocą = można przypisać do niej wartość używając operatora **:=**.

```go
package main

import (
	"fmt"
)

func main() {
	quantity := 10
	length, width := 10.5, 5.2
	customerName := "Damon Cole"

	fmt.Println("Quantity is", quantity)
	fmt.Println("Size is", length, "by", width)
	fmt.Println("Customer is", customerName)
}
```
Częste błędy: 
```go
quantity := 10
quantity := 5
```

Operator `:=` służy jednocześnie do deklaracji i inicjalizacji nowych zmiennych. Operator `:=` ten nie służy do modyfikacji już istniejących zmiennych. 

---------------
```go
quantity = 10
```
Jeśli zapomnimy `:` powyzsza instrukcja zostanie uznana za przypisanie a nie za deklaracje. Nie wolno przypisywac wartosci do niezadeklarowanych zmiennych. 

---------
```go
quantity := 10
quantity = "a"
```
Do danej zmiennej można przypisywac tylko wartości tego samego typu.

---------

Należy uzywać wszytkie zadeklarowane zmienne, które nie są używane. Jeśli usuniesz kod ktory uzywa zmiennej należy także wyeliminowac deklarację. 

### Reguły tworzenia nazw

Poniżej przedstawiono jak powinno tworzyć nazwy zmiennych, funkcji i typów w języku Go:

- nazwy powinny się zaczynać się od litery
- Jeśli zmienna rozpoczyna się **wielka literą** to dany element jest **eksportowany** i można go wykorzystywac **w innych pakietach** niż bieżacy. Dla przykładu `fmt.Println` stosowane jest wielkie `P`, dzięki czemu funkcji można używać w `main` i dowolnym innym pakiecie. Jeśli nazwa zmiennej zaczyna sie **małą literą** to oznacza że mozna z niego korzystac **tylko w tym bieżącym **pakiecie. 

Poprawne: 

```go
length
stack2
sales.Total
```
Niepoprawne
```go
2stack ← Nazwa **nie może** zaczynać się od cyfry
sales.Total ← **nie można** uzyć elementow z innych pakietow jeśli nazwa zaczyna się od małej litery
```

Powyzsze dwie cechy sa wymuszane przez język. Poniżej przedstawiono jeszcze kilka dodatkowych konwencji:

- Jeśli nazwa składa się z wielu słów, to słowo po pierwszym powinno zaczynać się wielką literą. Słowa powinno się łączyć bez stosowania spacji. Wykorzystujemy notację wielbłądzią **(camel notation)** Pierwsa litera powinna być wileka tylko wtedy gdy chcemy eksportować dany element poza pakiet. 
- dla oczywistych przykladów skraca się nazwę na przykład: `max` zamiast `maximum`, `i` zamiast `indeks`.

Poprawne:
```go
sheetLength
TotalUnits
i
```

Niezgodne z konwencjami:

```go
sheetlength ← kolejne słowa powinny zaczynać sie wielka literą
Total_Units ← nazwy powinny byc bezposrednio połączone
index ← zastąpic to skrótem 
```

Operacje matematyczne i porównania w języku Go wymagają używania wartości tego samego typu. Użycie wartości rożnych typów spowoduje błąd przy probie uruchomienia kodu. 

```go
var length float64 = 1.2
var width int = 2
fmt.Println("Powierczhnia wynosi", length*width)
fmt.Println("length > width", length > width)
```
`./ShortDeclaration.go:17:37: invalid operation: length * width (mismatched types float64 and int)
./ShortDeclaration.go:18:41: invalid operation: length > width (mismatched types float64 and int)`


Nowe wartosci ktore przypisujemy do zmiennych też muszą mieć ten sam typ.

```go
var length float64 = 1.2
var width int = 2
length = width
fmt.Println(length)
```

`command-line-arguments
./ShortDeclaration.go:17:11: cannot use width (variable of type int) as float64 value in assignment`

Rozwiązaniem powyzszych problemów jest zastosowanie **konwersji**, pozwalajacych na przekształcenie wartości jednego typu na inny typ. Wystarczy podać typ na który chcesz zmienić.

```go
var myInt int = 2
float64(myInt) 
```

- `float64` ← typ na który przeprowadzana jest konwersja,
- (myInt) ← przekształcana wartość.

Dokonanie konwersji:

```go
var length float64 = 1.2
var width int = 2
fmt.Println("Powierczhnia wynosi", length*width)
fmt.Println("length > width", length > width)
```
Wyniku czego otrzymamy:
`Powierczhnia wynosi 2.4
length > width? false`

Tak samo będzie tutaj: 

```go
var length float64 = 1.2
var width int = 2
length = float64(width)
fmt.Println(length)
```

Warto pamietać,że gdy zmieniamy wartość z typu `float64` na typ `int` część ułamkowa bedzie pomijana.

```go
var length float64 = 3.75
var width int = 5
width = int(length)
fmt.Println(width)
```
Otrzymamy: `3`

## Narzędzia jezyka Go

- `go build` ← kompiluje pliki z kodem do postaci plików binarnych
- `go run` ← kompiluje i uruchamia program bez zapisywania pliku wykonywalnego.  **Zaleca się korzystaniez niego dla małych programami**. Gdy wprowadzisz zmiany w kodzie, nie będziesz musiał wykonywać odrębnego kroku kompilacji.
- `go fmt` ← formatuje pliki źródłowe zgodnie ze standordowym formatowania języka Go
- `go version` ← wyświetla zainstalowaną wersję języka Go

Rozbudowane programy uruchamia sie po przez kolejne uzycie komend:

1) go fmt hello.go
2) go build hello.go
3) ./hello  


## Wywoływanie metod

W Go można definiować **metody**. Są to funkcje powiązane z wartościami określonego typu.  

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	var year int = now.Year()
	fmt.Println("Year is", year)
}
```

Pakiet `time` zawiera typ `Time`, który reprezentuje date(rok, miesiąc, dzień). Każda wartość typu `time.Time` udostępnia metodę `Year`, ktora zwraca rok. W kodzie ponizej metoda służy do wyświetlenia bieżącego roku.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	var year int = now.Year()
	fmt.Println("Year is", year)
}
```

Funkacja `time.Now` zwraca nową wartość typu `time.Time` z bieżącą datą i czasem. Wartość ta przypisywana jest do zmiennej `now`. Następnie wywoływana jest metoda `Year` dla wartości wskazanej przez `now`

`now.Year()` ← `now` ← Przechowuje wartość typu time.Time,
			`Year()` ← Wywołuje metodę Year dla wartości typu time.Time

Metoda `Year` zwraca liczbę całkowitą reprezentująca rok, który następnie jest wyświetlany.

Nastepny przyklad:

```go
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
```

Funkcja `strings.NewReplacer` przyjmuje argumenty w postaci zastępowanego podłańcucha `("#")` i podstawionego tekstu ("o"), a zwraca wartość typu `string.Replacer`. Gdy przekażesz łancuch znaków do metody `Replace` wartości typu `Replacer`, metoda zwroci zmodyfikowany łancuch znaków. 

Kropka oznacza że element podany po prawej stronie należy do jednostki podanej po lewej. 

!!! **Funkcje należą do pakietu, a metody do wartości.** !!!
## Obliczenie oceny

Stworzenie prostego programu w języku **go**, który po wprowadzeniu danych w procentach analizuje czy uczen zdał. Wynik poniżej 60% oznacza nie zdanie egzaminu. 

### Komentarze

W jezyku **go** komentarze oznacza się za pomocą dwóch ukośników **//**. Tak jak to jest w innch jezykach moż;liwe sa także *komentarze blokowe* oznaczane jako 

```go
/*
Treść komentarza
*/ 
```

### Pobieranie wyniku od użytkownika

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input := reader.ReadString('\n')
	fmt.Println(input)

}
```

W linijce 	*reader := bufio.NewReader(os.Stdin)* 
Wykorzystano funkcję *NewReader*, która pozwala na pobranie tekstu z klawiatury zawarte w strumieniu *os.Stdin*.
Następnie wywoływana jest metoda  *ReadString*, która zwraca dane wpisane przez użytkownika do runy nowego wiersza czyli do kliknięcia przycisku *"Enter"*. 

Po uruchomieniu ukarze się komunikat: 

*$ go run pass_fail.go 
/# command-line-arguments
./pass_fail.go:12:11: assignment mismatch: 1 variable but reader.ReadString returns 2 values*

Powyższy komunikat mówi że próbujemy zwrocić *dwie* wartości, a podana jest tylko *jedna* zmienna. 

Jezyk Go umożliwia zwracanie dowolnej liczby wartości przez funkcje i metody, najczęsciej dodatkowe zmienne informują o błędach. 

### Funkcje i metody zwracające wiele wartości

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input , err := reader.ReadString('\n')
	fmt.Println(input)
}
```

Dodanie do powyższego kody zmiennej **err** nie pozwoli na skompiluje, gdyż język go nie pozwala na skompilowanie kodu, gdzie znajduje się zmienna, której nie używamy.

### Zignorowanie wartosci błędu za pomocą pustego identyfikatora

Gdy nie chcemy używać zmniennej zwracanej przez funkcję lub metodę możemy uzyć **pustego identyfikatora** czyli **(_)**. 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input , _ := reader.ReadString('\n')
	fmt.Println(input)
}
```

### Obsługa błędu

Ignorowanie błędu nie jest dobrym pomysłem, gdyż jeśli by ten błąd się pojawił to nie zostalibysmy o tym poinformowani.

Najlepiej abysmy zostali poinformowani o błędzie i zatrzymali program. Do tego służy pakiet **log**, który zawiera funkcję **Fatal**. Pozwala ona na wwyświetlenie komunikatu w terminalu oraz zatrzymanie ("zabicia") programu. 

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)
}
```

*$ go run pass_fail.go 
Podaj wynik egzaminu: 50
2024/10/25 20:40:29 <nil>
exit status 1*


Funkcje takie jak ReedString zwracają wartość błędu *nil*, gdy błąd **nie nastąpił**. Chcemy aby program kończył jeśli zmienna `err` ma wartość inną od `nil`. Do tego służą funkcje warunkowe, które powodują wykonanie bloku kodu tylko wtedy, gdy spełniony jest dany warunek. 

### Instrukcje warunkowe 

Instrukcje warunkowe wykonuja instrukcje gdy spelniony zostanie warunek. Innymi słowy gdy warunek zwróci wartość wyniku `true`to zostanie wykonany kod w bloku warunkowym. Jeżeli wynik to `false` to blok warunkowy jest pomijany. 

```go
if 1 == 1 { // wynik true
	fmt.Println("Zostanę wyświetlony")
}
```

```go
if 1 >= 2 { //wynik false
	fmt.Println("A ja nie zostanę wyświtlony")
}
```

### Warunkowe rejestrowanie błędu krytycznego 

Powracając do naszego błędu, gdzie program kończył prace i zgłaszał błąd. 

*$ go run pass_fail.go 
Podaj wynik egzaminu: 50
2024/10/25 20:40:29 <nil>
exit status 1*

Wykorzystując instrukcje warunkowe do rozwiązania naszego problemu musimy przerwac program gdy wartość zmiennej `err` jest inna niż `nil`. Nasza instrukcja warunkowa wyglada następująco:

	if err != nil {
		log.Fatal(err)
	}


```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
```

Warto zaznaczyć że powyższy program wyświetla liczbę wpisaną przez użytkownika. 

### Zakrywanie nazw

Dobrą praktyką jest unikanie skrutów. Jednakże zmienna `error` jest złym pomysłem, gdyż zakryje zmienną typu *error*.


Podczas deklaracji zmiennej pamietaj, że nazwa zmiennej powinna być różna od istniejących zmiennych. Jeśli nie to w konsekwencji nowa zmienna zakryje poprzednią funkcjonalność. 


```go
package main

import (
	"fmt"
)

func main() {
	var int int =12
	var append string = "minut dadatkowych materiałów"
	var fmt string = "DVD"
	var count int 
	var languages = append([]string{}, "Espanol")
	fmt.Println(int, append, fmt, languages)
}
```

Powyższy kod nie skompiluje się ze względu na to zakrylisci nazwą zmiennej *typ zmiennej int*,  *funkcję append*, *nazwę pakietu*. 

Aby ten problem rozwiązać, wystarczy zmienić nazwy zmiennych:

```go

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

```

### Przekształcanie łancuchów znaków na liczby
Powracając do naszego programu `pass_fail.go`
Instrukcje warunkowe umożliwiają sprawdzanie podanego wyniku. Dodatkowo istnieje instrukcja `if/else` która określa zaliczenie albo jego brak.  

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if input >= 60 {
		status := "zaliczenie"
	} else {
		status := "brak zaliczenia"
	}
	fmt.Println("Status:", status)
}
```



*$ go run pass_fail.go 
# command-line-arguments
./pass_fail.go:18:14: invalid operation: input >= 60 (mismatched types string and untyped int)*

Przyczyna błędu jest wczytanie danej z klawiatury jako łańcucha znaków. Jezyk Go potrafi porównywać liczby tylko z innymi liczbami (dane które porownujemy w warunku muszą być tego samego typu), nie porównuje liczby z łancuchem znaków lub odwrotnie.

Nie można przekonwertować bezposrednio z typu `string` na `liczby`.

```go 
float64("2.6")
```
Wyświetli się wtedy taki komunikat:
*cannot convert "2.6" (type string) to type float64*

Nalezy zwrócić uwagę na:

- dane wprowadzone z klawiatury posiadają znak nowego wiersza **(\n)**, który trzeba usunąć
- resztę łańcucha znaków trzeba przekształcić na liczbę zmienno przecinkową
  

Aby usunąć znak nowego wiersza należy użyć funkcji `TrimSpace`, króra eliminuje wszystkie białe znaki(znak nowego wiersza, tabulacji, spacje) zarówno z początku jak i końca łańcucha znaków. 

```go
	s := "\t wczesniej był tabulator i nowy wiersze \n"
	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
```

```go
$ go run pass_fail.go 
         wczesniej był tabulator i nowy wiersze 

wczesniej był tabulator i nowy wiersze
```

`input = strings.TrimSpace(input)`

Wyniku czego pozostanie tylko wprowadzona liczba.

Teraz wykorzystujemy funkcję `ParseFloat` z pakietu `strconv` do przekształcenia liczby na wartosć typu float64.

```go
grade, err := strconv.ParseFloat(input,64)
```
Poowyzsza funkcja zwraca:
- wartość typu float64 `grade`
- ewentualny błąd `err` <- gdy nie da sie przekształcić
- łańcuch znaków `input`
- precyzja w bitach `64`

Po modyfikacjach kod wyglada następująco:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Podaj wynik egzaminu: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status := "zaliczenie"
	} else {
		status := "brak zaliczenia"
	}
	fmt.Println("Status:", status)
}
```

Niestety pojawia się następny błąd:


```go
	if grade >= 60 {
		status := "zaliczenie"
	} else {
		status := "brak zaliczenia"
	}
	fmt.Println("Status:", status)
```
```go
$ go run pass_fail.go 
# command-line-arguments
./pass_fail.go:28:3: declared and not used: status
./pass_fail.go:30:3: declared and not used: status
./pass_fail.go:32:25: undefined: status
```

Jka widać 2 pierwsze błędy dotyczą że zmienna została zadeklarowana a nie została użyta, a trzeci dotyczy że zmienna nie została zdefiniowana.

### Bloki i ich zasięg

Błąd narodził się ze względu na podzielenie kodu na **bloki**. Bloki mogą być zagnieżdzone jeden w drugim.  
Każda deklarowana zmienna ma zasięg, czyli zakres kodu w jakim jest widoczna. Próba użycia kończy się błędem. 






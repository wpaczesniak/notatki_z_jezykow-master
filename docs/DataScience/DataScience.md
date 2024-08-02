# Data Science

sprawdzenie typu zmiennej

```python
type(nazwa_zmiennej)
```

### zmienne

int - typ zmiennej całkowitej

float - typ zmiennej ułamkowej

str - typ zmiennej łańcuchowej (string) czyli text

bool - typ wartosć logiczna True/False

### Lista

Zbiór danych zapisanych pod nazwą zmiennej. W pythonie można zapisywać różne typy zmiennej do tej samej tablicy. Można zapisywać listy w listach.

```python
hall = 11.25
kit = 18.0
liv = 20.0
bed = 10.75
bath = 9.50

# House information as list of lists
house = [["hallway", hall],
         ["kitchen", kit],
         ["living room", liv],
         ["bedroom", bed],
         ["bathroom", bath]]

# Print out house
print(house)
```

### Indeksowanie elementów tablicy

Indeksowanie elementów tablicy zaczyna się od liczby 0.

```python
# Create the areas list
areas = ["hallway", 11.25, "kitchen", 18.0, "living room", 20.0,
"bedroom", 10.75, "bathroom", 9.50]

# Print out second element from areas
print(areas[1])

#Output
#11.25
```

Indeks -1 oznacza ostatni element w tablicy.

```python
# Print out last element from areas
print(areas[-1])

#Output
#9.5
```

Można wykorzystać przedział, jaki chcemy, aby został wyświetlony. Zwracany są elementy nieuwzględniającego ostatniego podanego indeksu.

```python

```

Można zwrócić elementy do konkretnego indeksu lub od konkretnego indeksu.

```python
#Użyj krojenia, aby utworzyć listę, downstairs, która zawiera 6 pierwszych
#elementów obszarów.

downstairs = areas[:6]

#Output
#['hallway', 11.25, 'kitchen', 18.0, 'living room', 20.0]

#Utwórz górę, jako ostatnie 4 elementy obszarów. Tym razem uprość cięcie, pomijając
#indeks końcowy.

upstairs = areas[6:]

#Output
#['bedroom', 10.75, 'bathroom', 9.5]
```

Podzbiory

Lista Pythona może również zawierać inne listy. Do podzbioru list list można użyć tej samej techniki, co poprzednio: nawiasów kwadratowych. Wyglądałoby to mniej więcej tak dla listy, house:

```python
#Subset the house list to get the float 9.5.

house = [["hallway", 11.25],
         ["kitchen", 18.0],
         ["living room", 20.0],
         ["bedroom", 10.75],
         ["bathroom", 9.50]]

# Subset the house list
house[-1][1]

#Output
#9.5
```

Zmiana wartości w liście

Można dokonać zmiany wartości w liście przypisując jej jakąś wartość.

```python
areas = ["hallway", 11.25, "kitchen", 18.0, "living room", 20.0, "bedroom", 10.75,
"bathroom", 9.50]

areas[-1] = 10.50

areas[4] = "chill zone"
```

Można także przypisać wartości do kilku wybranych indeksów.

```python
areas[2:4] = ["garage", 14.45]
```

Można także dopisać nowe elementy do istniejącej tablicy

```python
areas_1 = areas + ["poolhouse", 24.5]
```

Można także usuwać elementy z listy.

```python
del areas[10:12]

#Output
# 'hallway', 11.25, 'kitchen', 18.0, 'chill zone', 20.0, 'bedroom', 10.75,
# 'bathroom', 10.5, 'garage', 15.45]
```

Jak jest zmienna to ona przechowuje wskaźnik na listę.

Zapisanie tablicy pod nową zmienną **y**

```python
x = ["a", "b", "c"]
y = x
y[1] = "z"
y

#Output
#['a', 'z','c']

x
#Output
#['a', 'z','c']
```

Spowoduje skopiowanie odniesienia do listy, a nie same wartości. Wskazuje ona na to samo miejsce w pamięci

Aby wskazać inne miejsce w pamięci należy użyć funkcji listy.

```python
x = ["a", "b", "c"]
y = list(x)
y = x[:]
y[1] = "z"
x

#Output
#['a', 'b','c']

x
#Output
#['a', 'z','c']
```

### Funkcje

Fragment kodu, który można wykorzystywać wielokrotnie.

Agregacje

max() - znajduje maksymalna wartość

min() - znajduje minimalną wartość

round() - zaokrągla

len() - zwraca długość tablicy

append() - dodaje wartość na końcu listy

remove() - usuwa pierwszy argument z listy

reverse() - odwraca listę

index() - zwraca wartość indeksu, na którym znajduje się dany element

count() - zlicza wystąpienia danego argumentu

wgląd do dokumentacji

```python
help(nazwa_funkcji)

# LUB

?nazwa_funkcji
```

Zwracanie indeksu

```python
areas = ["hallway", 11.25, "kitchen", 18.0, "living room", 20.0, "bedroom", 10.75,
"bathroom", 9.50]

areas.index("kitchen")
#Output
#2
```

Zliczenie ile razy w tablicy pojawił się dany argument

```python
areas.count(11.25)

#Output
#1
```

Zmiana pierwszej litery na dużą literę

```python
sister
#Output
#'liz'

sister.capitalize()
#Output
#'Liz'
```

Zmiana niektórych części w ciągu innymi częściami

```python
sister.replace("z", "sa")
#Output
#lisa
```

Metody mogą się zachowywać inaczej dla innego rodzaju obiektu.

### Pakiet NumPy

//TO DO Opisać to bardziej

Służy do wykonywania skomplikowanych obliczeń w Pythonie. Wykonywanie obliczeń na zwykłej liście było bardzo nieefektywne i było bardzo czasochłonne i bardzo skomplikowane.

```python
#instalacja
pip3 install numpy
```

Różnice między używaniem NumPy oraz nie.

```python
height = [1.73, 1.68, 1.71, 1.89, 1.79]
weight = [65.4, 59.2, 63.6, 88.4, 68.7]
weight / height ** 2

#Output
#TypeError: unsupported operand type(s) for ** or pow(): 'list' and 'int'

```

Wykorzystanie NumPy

```python
np_height = np.array(height)
np_weight = np.array(weight)
np_weight / np_height ** 2

#Output
#array([21.85171573, 20.97505669, 21.75028214, 24.7473475 , 21.44127836])
```

Różnice w dodawaniu tablic

```python
python_list = [1, 2, 3]
numpy_array = np.array([1, 2, 3])

python_list + python_list
#Output
# [1, 2, 3, 1, 2, 3]

numpy_array + numpy_array
#Output
#array([2, 4, 6])
```

```python
np.array([True, 1, 2]) + np.array([3, 4, False])

#Output
#np.array([4, 3, 0]) + np.array([0, 2, 2])
```

Wyświetlanie wato sci które spełniaja warunek

```python
bmi = array([21.85171573, 20.97505669, 21.75028214, 24.7473475, 21.44127836])
```

Dostęp do elementu

```python
Kod: bmi[1]

#Output
#Wynik: 20.975
```

Warunkowe filtrowanie

```python
Kod: bmi > 23

#Output
#array([False, False, False, True, False])
```

Elementy spełniające warunek

```python
bmi[bmi > 23]

#Output
#array([24.7473475])
```

### Dwuwymiarowe tablice

Określenie ile tablica ma wierszy i kolumn wykorzystuje się metodę **shape**

Zwracanie elementu w tablicy dwuwymiarowej

```python
np_2d[0][2]

#LUB

np_2d[0, 2]

#Wyświetlenie tylko drugiej i trzeciej kolumny
np_2d[:, 1:3]

#Piewszy argument to wiersz a drugi to kolumna
```

### Intermediate Python

### Matplot

lineplot

```python
import matplotlib.pyplot as plt

year = [1950, 1970, 1990, 2010]
pop = [2.519, 3.692, 5.263, 6.972]

plt.plot(year, pop)
plt.show()

```

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled.png)

**Scatter Plot**

```python
import matplotlib.pyplot as plt

year = [1950, 1970, 1990, 2010]
pop = [2.519, 3.692, 5.263, 6.972]

plt.scatter(year, pop)
plt.show()
```

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%201.png)

Skala logarytmiczna

```python
plt.xscale('log')
```

### Histogramy

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%202.png)

Jeśli potrzebujemy pomocy z funkcją i znaleźć o niej informację należy skorzystać z poniższej komendy. Należy pamiętać, że należy o skrócie biblioteki.

```python
import matplotlib.pyplot as plt

help(plt.hist)
```

Tworzenie Histogramów w matlib

```python
values = [0, 0.6, 1.4, 1.6, 2.2, 2.5, 2.6, 3.2, 3.5, 3.9, 4.2, 6]

plt.hist(values, bins=3) #Określa ile ma być podziałów na wykresie
plt.show()
```

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%203.png)

`clf()` tworzy nowy czysty wykres tak by dane przedstawić na osobnym wykresie.

`xlabel()` Umożliwiają dodanie opisu osi x

`ylabel()` Umożliwiają dodanie opisu osi y

`title()` Dodanie tytuły dla wykresu

`yticks()` Dodaje te wartości na wykresie na osi y ktore chcemy wyswietlic

Możemy przypisać nazwy do tych wartości. Należy pamiętać o tym, aby ta lista była tej samej długości co poprzednia.

```python
plt.yticks([0, 2, 4, 6, 8, 10], ['0', '2B', '4B', '6B', '8B', '10B'])
```

Można dodawać nowe dane do wykresu

```python
import matplotlib.pyplot as plt

# Dane z lat 1950-2100, przykładowe wartości dla pierwszych i ostatnich lat
year = [1950, 1951, 1952, ..., 2100]
pop = [2.538, 2.57, 2.62, ..., 10.85]

# Dodanie dodatkowych danych z wcześniejszych lat
year = [1800, 1850, 1900] + year
pop = [1.0, 1.262, 1.650] + pop

plt.plot(year, pop)
plt.xlabel('Year')
plt.ylabel('Population')
plt.title('World Population Projections')
plt.yticks([0, 2, 4, 6, 8, 10], ['0', '2B', '4B', '6B', '8B', '10B'])
plt.show()

```

Zmiana rozmiaru punktów (Ich powiększenie)

```python
plt.scatter(gdp_cap, life_exp, s = np_pop)
```

Pozwala na zmianę rozmiaru punktow by staly się one zależne od swojego rozmiaru

Caly kod()

```python
# Import numpy as np
import numpy as np

# Store pop as a numpy array: np_pop
np_pop = np.array(pop)

# Double np_pop
np_pop = np_pop * 2

# Update: set s argument to np_pop
plt.scatter(gdp_cap, life_exp, s = np_pop)

# Previous customizations
plt.xscale('log')
plt.xlabel('GDP per Capita [in USD]')
plt.ylabel('Life Expectancy [in years]')
plt.title('World Development in 2007')
plt.xticks([1000, 10000, 100000],['1k', '10k', '100k'])

# Display the plot
plt.show()
```

Rezultat

![Untitled](DataScience/Untitled4.png)

Argument `c` w funkcji `scatter` oznacza dostosowanie kolorów.

Można dodać opisy do poszczególnych punktów na wykresie

```python
# Scatter plot
plt.scatter(x = gdp_cap, y = life_exp, s = np.array(pop) * 2, c = col, alpha = 0.8)

# Previous customizations
plt.xscale('log')
plt.xlabel('GDP per Capita [in USD]')
plt.ylabel('Life Expectancy [in years]')
plt.title('World Development in 2007')
plt.xticks([1000,10000,100000], ['1k','10k','100k'])

# Additional customizations
plt.text(1550, 71, 'India')
plt.text(5700, 80, 'China')

# Add grid() call
plt.grid(True)

# Show the plot
plt.show()
```

![Untitled](DataScience/Untitled5.png)

### Słowniki

Jeśli korzystamy z list to musimy korzystać z indeksów co jest mało intuicyjne oraz mało wygodne xD.

Lepsze do tego jest wykorzystanie słowników.

Poniżej przedstawiono porównanie wykorzystania słowników i list

```python
pop = [30.55, 2.77, 39.21]
countries = ["afghanistan", "albania", "algeria"]
ind_alb = countries.index("albania")

# Wynik dla ind_alb
print(ind_alb)  # Wynik: 1

# Dostęp do populacji Albanii używając uzyskanego indeksu
print(pop[ind_alb])  # Wynik: 2.77
```

Słownik

```python
pop = [30.55, 2.77, 39.21]
countries = ["afghanistan", "albania", "algeria"]

world = {"afghanistan": 30.55, "albania": 2.77, "algeria": 39.21}
print(world["albania"])  # Wynik: 2.77
```

Słownik skonstruowany jest w sposób **_klucz : wartość_** Pozwala to uniknąć . Dostęp do wartości obywa się poprzez klucz.

Kolejny przykład stworzenia słownika (Przykład pochodzi z Data Campa)

```python
# Definition of countries and capital
countries = ['spain', 'france', 'germany', 'norway']
capitals = ['madrid', 'paris', 'berlin', 'oslo']

# From string in countries and capitals, create dictionary europe
europe = {'spain':'madrid', 'france':'paris', 'germany':'berlin', 'norway':'oslo' }

# Print europe
print(europe)

# output:
# {'spain': 'madrid', 'france': 'paris', 'germany': 'berlin', 'norway': 'oslo'}
```

Aby wyświetlić same klucze, należy użyć metody `keys()`

```python
# Definition of dictionary
europe = {'spain':'madrid', 'france':'paris', 'germany':'berlin', 'norway':'oslo' }

# Print out the keys in europe
print(europe.keys())

# Print out value that belongs to key 'norway'
print(europe['norway'])
```

Klucze w słowniku są obiektami, które są niezmienne. Klucze powinny być unikalne. Możemy zmieniać jedynie wartości tych kluczy.

Można dodawać nowe rekordy słownika.

wykonujemy to poprzez

```python
world['sealand'] = 0.000027
```

Aby sprawdzić, czy dodanie przebiegło poprawnie, możemy sprawdzić, czy w słowniku znajduje się ten rekord:

```python
"sealand" in world
#True
```

Można zmieniać wartości w słowniku

```python
world['sealand'] = 0.000029
```

Można usuwać elementy w słowniku poprzez wykonanie komendy `del`

```python
del(world['sealand'])
```

Różnice między listą a słownikiem

**Lista:**

- Możliwość wyboru, aktualizacji i usuwania elementów przy użyciu nawiasów kwadratowych `[]`.
- Indeksowana zakresem numerów.
- Kolekcja wartości, gdzie kolejność ma znaczenie, umożliwiająca wybór całych podzbiorów.

**Słownik:**

- Możliwość wyboru, aktualizacji i usuwania elementów przy użyciu nawiasów kwadratowych `[]`.
- Indeksowany unikalnymi kluczami.
- Tabela wyszukiwania z unikalnymi kluczami.

Istnieje możliwość stworzenia słownika, który zawiera w sobie inne mniejsze słowniki.

Poniżej widać przyklad z DataCampa

```python
# Dictionary of dictionaries
europe = { 'spain': { 'capital':'madrid', 'population':46.77 },
           'france': { 'capital':'paris', 'population':66.03 },
           'germany': { 'capital':'berlin', 'population':80.62 },
           'norway': { 'capital':'oslo', 'population':5.084 } }

# Print out the capital of France
print(europe['france']['capital'])

# Create sub-dictionary data
data = {'capital':'rome', 'population':59.83}

# Add data to europe under key 'italy'
europe['italy'] = data

# Print europe
print(europe)
```

### Pandas podstawy

tablice NumPy nie radzą sobie dobrze z różnymi rodzajami danych.

W pandasie przechowujemy dane tabelaryczne. Przechowywane są w ramce danych DataFrame

Aby utworzyć ramkę danych, można skorzystać ze słownika.

```python
dict = {
    "country": ["Brazil", "Russia", "India", "China", "South Africa"],
    "capital": ["Brasilia", "Moscow", "New Delhi", "Beijing", "Pretoria"],
    "area": [8.516, 17.10, 3.286, 9.597, 1.221],
    "population": [200.4, 143.5, 1252, 1357, 52.98]
}
```

```python
import pandas as pd

brics = pd.DataFrame(dict)
```

Dodanie nowych indeksów można zrobić poprzez:

```python
brics.index = ["BR", "RU", "IN", "CH", "SA"]
brics
```

Importowanie danych z pliku csv odbywa się przy użyciu komendy

```python
brics = pd.read_csv("path/to/brics.csv")
brics
```

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%206.png)

Aby ustalić indeksy w ramce danych korzystamy z argumentu index_col.

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%207.png)

Kolejny przykład stworzenia słownika, który składa się w sobie z innych słowników.

```python
# Pre-defined lists
names = ['United States', 'Australia', 'Japan', 'India', 'Russia', 'Morocco', 'Egypt']
dr =  [True, False, False, False, True, True, True]
cpc = [809, 731, 588, 18, 200, 70, 45]

# Import pandas as pd
import pandas as pd

# Create dictionary my_dict with three key:value pairs: my_dict
my_dict = {'country' : names, 'drives_right' : dr, 'cars_per_cap' : cpc}

# Build a DataFrame cars from my_dict: cars
cars = pd.DataFrame(my_dict)

# Print cars
print(cars)

#output:
#             country  drives_right  cars_per_cap
#    0  United States          True           809
#    1      Australia         False           731
#    2          Japan         False           588
#    3          India         False            18
#    4         Russia          True           200
#    5        Morocco          True            70
#    6          Egypt          True            45
```

Ustawienie własnych indeksów wykonuje się poprzez użycie komendy.

```python
# Definition of row_labels
row_labels = ['US', 'AUS', 'JPN', 'IN', 'RU', 'MOR', 'EG']

# Specify row labels of cars
cars.index = row_labels
```

**Dostep do danych**

Aby wyświetlić tylko jedną kolumne można to zrobić na dwa sposoby

```python
brigs["country"]

#LUB

brigs.country
```

Aby zachować dane tego samego typu jak w początkowej ramce danych, należy skorzystać z podwójnych nawiasów kwadratowych. W przeciwnym razie dane zostaną zapisane w serii.

```python
brigs[["country"]]
```

Można wybrać typ ramki danych dwie kolumny lub więcej.

```python
brigs[["country", "capital"]]
```

/TO DO
Doczytać o tym więcej

Aby uzyskać dostęp do wybranych danych, należy skorzystać z indeksów. Indeksy numerowane są od 0.

```python
brigs[1:4]
```

Dostęp do określonego rzedu, wierszu otrzymujemy poprzez użycie komendy `loc`. Pojedynczy nawias kwadratowy gdy dane są typu Serii.

```python
brigs.loc["RU"]
```

Podwójny nawias kwadratowy powoduje ze dane są w typu ramki danych.

```python
brigs.loc[["RU"]]
```

Typ Series używany gdy chcesz używać danych jako operacja typowe dla sekwencji, czyli takie jak obliczanie statystyk, funkcje agregujące. Typ DataFrame podwójne nawiasy używane są gdy chcemy użyć filtrowania, dołączania kolumn czy tabel.

Można uzyskać dostęp do wielu wierszy.

```python
brigs.loc[["RU","IN","CH"]]
```

Można uzyskać dostęp również wybrane wiersze jak i kolumny.

```python
brigs.loc[["RU", "IN", "CH"], ["country", "capital"]]
```

Aby wybrać tylko dwie kolumny z użyciem `loc`. Należy skorzystać z następującej składni.

```python
brigs.loc[:, ["country", "capital"]]
```

**Podsumowanie**

- **Użycie nawiasów kwadratowych**:
  - Dostęp do kolumn: `brics[["country", "capital"]]`
  - Dostęp do wierszy poprzez wycinanie: `brics[1:4]`
- **Użycie `loc` (dostęp oparty na etykietach)**:
  - Dostęp do wierszy: `brics.loc[["RU", "IN", "CH"]]`
  - Dostęp do kolumn: `brics.loc[:, ["country", "capital"]]`
  - Dostęp do wierszy i kolumn jednocześnie: `brics.loc[["RU", "IN", "CH"], ["country", "capital"]]`
  Można uzyskać dostęp przy użyciu indeksu. Poniższe wyrażenie:
  ```python
  brigs.loc[["RU", "IN", "CH"], ["country", "capital"]]
  ```
  Można zastąpić indeksami z wykorzystaniem metody `iloc` .
  ```python
  brigs.iloc[[1,2,3], [0,1]]
  ```
  Analogicznie
  ```python
  brigs.loc[:, ["country", "capital"]]
  ```
  ```python
  brigs.iloc[:, [0,1]]
  ```
  ### Operatory logiczne
  - `<` - Strictly less than
  - `<=` - Less than or equal
  - `>` - Strictly greater than
  - `>=` - Greater than or equal
  - `==` - Equal
  - `!=` - Not equal
  `AND`
  `OR`
  `NOT`
  Porównanie array w NumPy przy użyciu `and`, `or`, `not` wyświetla błąd.
  ```python
  bmi = array([21.85171573, 20.97505669, 21.75028214, 24.7473475, 21.44127836])

  bmi > 21 and bmi < 22

  #Output
  #ValueError: The truth value of an array with more than one element is ambiguous.
  #Use a.any() or a.all()
  ```
  Należy użyć odpowiadających metod:
  ```python
  logical_and()
  logical_or()
  logical_not()
  ```
  Poprawne użycie wygląda następująco:
  ```python
  np.logical_and(bmi > 21, bmi < 22)

  #Output
  #array([True, False, True, False, True], dtype=bool)
  ```
  Wyświetlenie odpowiednich wartości spełniających warunki.
  ```python
  bmi[np.logical_and(bmi > 21, bmi < 22)]

  #Output
  #array([21.852, 21.75, 21.441])

  ```
  ### Wyrażenie warunkowe
  ```python
  # Define variables
  room = "bed"
  area = 14.0

  # if-elif-else construct for room
  if room == "kit" :
      print("looking around in the kitchen.")
  elif room == "bed":
      print("looking around in the bedroom.")
  else :
      print("looking around elsewhere.")

  # if-elif-else construct for area
  if area > 15 :
      print("big place!")
  elif area > 10:
      print("medium size, nice!")
  else :
      print("pretty small.")
  ```
  Funkcja `elif` działa wtedy gdy warunek pierwszy oznaczony jako `if` nie został spełniony. Jeśli pierwszy warunek został spełniony to operacja `elif` nie zostanie wykonana. Może być wiele warunków `elif` . Jeśli pierwszy warunek `if` nie zostanie spełniony to przechodzi po kolei po kolejnych warunkach `elif`.
  ### Filtrowanie danych
  Jeśli w ramce danych będzie kolumna, której będą wartości logiczne `True`, `False` , to nie trzeba pisać dodatkowego warunku.
  ```python
  # Import cars data
  import pandas as pd
  cars = pd.read_csv('cars.csv', index_col = 0)
  print(cars)
  #				 cars_per_cap        country  drives_right
  #    US            809  United States          True
  #    AUS           731      Australia         False
  #    JPN           588          Japan         False
  #    IN             18          India         False
  #    RU            200         Russia          True
  #    MOR            70        Morocco          True
  #    EG             45          Egypt          True

  # Extract drives_right column as Series: dr
  dr = cars["drives_right"]

  # Use dr to subset cars: sel
  sel = cars[dr]

  # Print sel
  print(sel)

  #Output
  #         cars_per_cap        country  drives_right
  #    US            809  United States          True
  #    RU            200         Russia          True
  #    MOR            70        Morocco          True
  #    EG             45          Egypt          True

  ```
  Można zapisać poniższe komendy krocej
  ```python
  dr = cars["drives_right"]

  sel = cars[dr]
  ```
  ```python
  sel = cars[cars['drives_right']]
  ```
  Gdy należy spełnić wiele warunków, warto skorzystać ze wcześniejszych funkcji takich jak [**`np.logical_and()`**](http://docs.scipy.org/doc/numpy-1.10.0/reference/generated/numpy.logical_and.html), [**`np.logical_or()`**](http://docs.scipy.org/doc/numpy-1.10.0/reference/generated/numpy.logical_or.html) lub [**`np.logical_not()`**](http://docs.scipy.org/doc/numpy-1.10.0/reference/generated/numpy.logical_not.html),
  ```python
  # Import cars data
  import pandas as pd
  cars = pd.read_csv('cars.csv', index_col = 0)

  # Import numpy, you'll need this
  import numpy as np

  # Create medium: observations with cars_per_cap between 100 and 500

  cpc = cars['cars_per_cap']
  between = np.logical_and(cpc > 100, cpc < 500)
  medium = cars[between]

  # Print medium
  print(medium)
  ```

### Pęlte while i for.

Pętla `while` za każdym razem sprawdza warunek. Pętla `for` wykonuje określoną liczbę iteracji.

```python
# Initialize offset
offset = -6

# Code the while loop
while offset != 0 :
    print("correcting...")
    if offset > 0 :
      offset -= 1
    else :
      offset += 1
    print(offset)
```

**Pętla for**

```python
fam = [1.73, 1.68, 1.71, 1.89]
for height in fam:
    print(height)

#Output
#1.73
#1.68
#1.71
#1.89
```

metoda `enumerate` pozwala na uzyskanie indeksu.

```python
fam = [1.73, 1.68, 1.71, 1.89]
for index, height in enumerate(fam):
    print("index " + str(index) + ": " + str(height))

#Output
#index 0: 1.73
#index 1: 1.68
#index 2: 1.71
#index 3: 1.89
```

Występuje funkcja `capitalize()` , która wyświetla pojedyncze znaki ze słowa.

```python
for c in "family":
    print(c.capitalize())

#Output
#F
#A
#M
#I
#L
#Y

```

```python
# house list of lists
house = [["hallway", 11.25],
         ["kitchen", 18.0],
         ["living room", 20.0],
         ["bedroom", 10.75],
         ["bathroom", 9.50]]

# Build a for loop from scratch
for name, area in house:
    print("the " + str(name) + " is " + str(area) + " sqm" )

    #output:
    #the hallway is 11.25 sqm
    #the kitchen is 18.0 sqm
    #the living room is 20.0 sqm
    #the bedroom is 10.75 sqm
    #the bathroom is 9.5 sqm
```

Pętla w słowniku

```python
world = {"afghanistan": 30.55, "albania": 2.77, "algeria": 39.21}
for key, value in world:
    print(key + " -- " + str(value))

#Output
#ValueError: too many values to unpack (expected 2)
```

Ten błąd wynika z tego, że podczas iteracji po słowniku `world` bezpośrednio, otrzymuje się tylko klucze. Aby uzyskać pary klucz-wartość, należy użyć metody `.items()`:

```python
for key, value in world.items():
    print(key + " -- " + str(value))

#Output
#Algeria -- 39.21
#Afghanistan -- 30.55
#Albania -- 2.77
```

Trzeba pamiętać, że kolejność ma znaczenie. **Pierwsza zmienna otrzymuje klucz, a druga wartość.**

**Iterowanie po dwóch tablicach.**

```python
import numpy as np

np_height = np.array([1.73, 1.68, 1.71, 1.89, 1.79])
np_weight = np.array([65.4, 59.2, 63.6, 88.4, 68.7])
meas = np.array([np_height, np_weight])
for val in meas :
	print(val)

#Output
#[1.73, 1.68, 1.71, 1.89, 1.79]
#[65.4, 59.2, 63.6, 88.4, 68.7]
```

W powyższym przypadku pętla for wypisuje cala tablice w każdej iteracji.

Wypisanie każdego elementu tablicy wykonuje się poprzez wykorzystanie funkcji `nditer()`.

```python
import numpy as np

np_height = np.array([1.73, 1.68, 1.71, 1.89, 1.79])
np_weight = np.array([65.4, 59.2, 63.6, 88.4, 68.7])
meas = np.array([np_height, np_weight])

for val in np.nditer(meas):
    print(val)

#Output
#1.73
#1.68
#1.71
#1.89
#1.79
#65.4
#...

```

**Kolejny przykład z Data Campa:**

`np_height,` tablicę NumPy zawierającą wysokość zawodników Major League Baseball, oraz `np_baseball`, dwuwymiarową tablicę NumPy zawierającą zarówno `wysokość` (pierwsza kolumna), jak i `wagę` (druga kolumna) tych zawodników.

```python
# Import numpy as np
import numpy as np

# For loop over np_height
for x in np_height :
    print(str(x) + " inches")

# For loop over np_baseball
for x in np.nditer(np_baseball) :
    print(x)
```

### Iterowanie po DataFrame za pomocą for

DataFrame działają inaczej niż 2D NumPy. Nie wyświetla każdego wiersza tylko wyświetli nazwy etykiet czyli nazwy kolumn.

```python
import pandas as pd
brics = pd.read_csv("brics.csv", index_col = 0)
for val in brics:
	print(val)
```

W takim przypadku otrzymamy nazwy kolumn

`column,`

`capital,`

`area,`

`population`

### iterrows()

Metoda `iterrows()` w bibliotece pandas w Pythonie służy do iterowania po wierszach DataFrame. Dla każdej iteracji zwraca parę (indeks, seria), gdzie "indeks" to etykieta wiersza, a "seria" to obiekt Series zawierający dane danego wiersza.

```python
import pandas as pd
brics = pd.read_csv("brics.csv", index_col = 0)
for lab, row in brics.iterrows():
	print(lab)
	print(row)
```

**Brazylia (BR):**

- **Kraj:** Brazil
- **Stolica:** Brasilia
- **Powierzchnia:** 8.516 (jednostka nie podana, prawdopodobnie mln km²)
- **Populacja:** 200.4 (jednostka nie podana, prawdopodobnie mln)

**Rosja (RU):**

- **Kraj:** Russia
- **Stolica:** Moscow
- **Powierzchnia:** 17.1 (jednostka nie podana, prawdopodobnie mln km²)
- **Populacja:** 143.5 (jednostka nie podana, prawdopodobnie mln)

**Można wyświetlić tylko stolice danego indeksu.**

```python
import pandas as pd
brics = pd.read_csv("brics.csv", index_col = 0)
for lab, row in brics.iterrows():
	print(lab + ": " + row["capital"])

#Output
#BR: Brasilia
#RU: Moscow
#IN: New Delhi
#CH: Beijing
#SA: Pretoria
```

Można dodawać nową kolumnę do DataFrame.

```python
import pandas as pd
brics = pd.read_csv("brics.csv", index_col = 0)
for lab, row in brics.iterrows():
	brics.loc[lab, "name_length"] = len(row["country"])
print(brics)

#Output
#| Code | Country       | Capital   | Area   | Population | Name Length |
#|------|---------------|-----------|--------|------------|-------------|
#| BR   | Brazil        | Brasilia  | 8.516  | 200.40     | 6           |
#| RU   | Russia        | Moscow    | 17.100 | 143.50     | 6           |
#| IN   | India         | New Delhi | 3.286  | 1252.00    | 5           |
#| CH   | China         | Beijing   | 9.597  | 1357.00    | 5           |
#| SA   | South Africa  | Pretoria  | 1.221  | 52.98      | 12          |
```

Tutaj może nastąpic spadek wydajności. Należy całą kolumnę DataFrame

```python
import pandas as pd
brics = pd.read_csv("brics.csv", index_col = 0)
brics["name_length"] = brics["country"].apply(len)
print(brics)

#output
#Code| Country       | Capital   | Area   | Population | Name Length |
#BR  | Brazil        | Brasilia  | 8.516  | 200.40     | 6
#RU  | Russia        | Moscow    | 17.100 | 143.50     | 6
#IN  | India         | New Delhi | 3.286  | 1252.00    | 5
#CH  | China         | Beijing   | 9.597  | 1357.00    | 5
#SA  | South Africa  | Pretoria  | 1.221  | 52.98      | 12

```

## Joining Data with Pandas

Informacje o danych z pliku

```python
wards = pd.read_csv("Ward_Offices.csv")
print(wards.head())
print(wards.shape)

#Output
#ward	alderman	          address	        zip
#1	    Proco "Joe" ...	  2058 NORTH W...	60647
#2	    Brian Hopkins	    1400 NORTH ...	60622
#3	    Pat Dowell	      5046 SOUTH S...	60609
#4	    William D. B...	  435 EAST 35T...	60616
#5	    Leslie A. Ha...	  2325 EAST 71...	60649
#(50,4)
```

Drugi plik

```python
census = pd.read_csv("Ward_Census.csv")
print(census.head())
print(census.shape)

#Output
#ward	  pop_2000	  pop_2010	change	address	           zip
#   1	  52951	      56149	    6%	    2765 WEST SA...	   60647
#   2	  54361	      55805	    3%	    WM WASTE MAN...	   60622
#   3	  40385	      53039	    31%	    17 EAST 38TH...	   60653
#   4	  51953	      54589	    5%	    31ST ST HARB...	   60653
#   5	  55302	      51455	   -7%	    JACKSON PARK...	   60637
```

Marging tables

**Inner join**

Poniższy kod pobiera pierwszą ramkę DataFrame i łączy ją z drugą ramką DataFrame. i łączy ją po kolumnie **_ward._**

```python
wards_census = wards.merge(census, on="ward")
print(wards_census.head(4))

#Output
#ward	  alderman	        address_x	         zip_x	  pop_2000	 pop_2010	  change	  address_y	        zip_y
#   1	  Proco "Joe" ...	  2058 NORTH W...	   60647	  52951	     56149	    6%	      2765 WEST SA...	  60647
#   2	  Brian Hopkins	    1400 NORTH ...	   60622	  54361	     55805	    3%	      WM WASTE MAN...	  60622
#   3	  Pat Dowell	      5046 SOUTH S...	   60609	  40385	     53039	    31%	      17 EAST 38TH...	  60653
#   4	  William D. B...	  435 EAST 35T...	   60616	  51953	     54589	    5%	      31ST ST HARB...	  60653

```

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%208.png)

Warto pamiętać ze jeśli tabele mają takie same nazwy kolumn to Pandas automatycznie doda przyrostki **_\_x_** lub **_\_y._**

```python
print(wards_census.columns)

#Index(['ward', 'alderman', 'address_x', 'zip_x', 'pop_2000', 'pop_2010', 'change',
# 'address_y', 'zip_y'],
# dtype='object')
```

Możemy to kontrolować dodając argument `suffixes` do którego będa przypisane wartosci w krotce. W poniższy przypadku będzie to oznaczalo że **_\_ward_** zostanie przypisany do kolumn należących wcześniej do pierwszej tabeli a **_\_cen_** zostanie przypisany do kolumn należących wcześniej do drugiej tabeli.

```python
wards_census = wards.merge(census, on='ward', suffixes=('_ward', '_cen'))
print(wards_census.head())
print(wards_census.shape)

#Output
#ward	  alderman	         address_ward	      zip_ward	  pop_2000	  pop_2010	  change	   address_cen	    zip_cen
#   1	  Proco "Joe" ...	   2058 NORTH W...	  60647	      52951	      56149	      6%	       2765 WEST SA...	60647
#   2	  Brian Hopkins	     1400 NORTH ...	    60622	      54361	      55805	      3%	       WM WASTE MAN...	60622
#   3	  Pat Dowell	       5046 SOUTH S...	  60609	      40385	      53039	      31%	       17 EAST 38TH...	60653
#   4	  William D. B...	   435 EAST 35T...	  60616	      51953	      54589	      5%	       31ST ST HARB...	60653
#   5	  Leslie A. Ha...	   2325 EAST 71...	  60649	      55302	      51455	      -7%	       JACKSON PARK...	60637

```

Aby policzyć ile było wystąpień danego rekordu w danej tabeli używamy metody `value_counts()`

```python
# Merge the taxi_owners and taxi_veh tables setting a suffix
taxi_own_veh = taxi_owners.merge(taxi_veh, on='vid', suffixes=('_own','_veh'))

# Print the value_counts to find the most popular fuel_type
print(taxi_own_veh['fuel_type'].value_counts())

#Output
#HYBRID                    2792
#GASOLINE                   611
#FLEX FUEL                   89
#COMPRESSED NATURAL GAS      27
#Name: fuel_type, dtype: int64
```

### Relacje w tabeli

**Relacja jeden do jeden**

W tej relacji każdy wiersz w lewej tabeli jest powiązany z tylko jednym wierszem w prawej tabeli.

**Relacja jeden do wielu**

W tej relacji każdy wiersz po lewej tabeli jest powiazany z jednym lub większą liczbą wierszy w prawej tabeli. Na przykład w jednym oddziale pracuje wiele firm.

```python
ward_licenses = wards.merge(licenses, on='ward', suffixes=('_ward', '_lic'))
print(ward_licenses.head())
```

| ward | alderman        | address_ward    | zip_ward | account | aid | business        | address_lic     |
| ---- | --------------- | --------------- | -------- | ------- | --- | --------------- | --------------- |
| 1    | Proco "Joe" ... | 2058 NORTH W... | 60647    | 12024   | nan | DIGILOG ELEC... | 1038 N ASHLA... |
| 1    | Proco "Joe" ... | 2058 NORTH W... | 60647    | 14446   | 743 | EMPTY BOTTLE... | 1035 N WESTE... |
| 1    | Proco "Joe" ... | 2058 NORTH W... | 60647    | 14624   | 775 | LITTLE MEL'S... | 2205 N CALIF... |
| 1    | Proco "Joe" ... | 2058 NORTH W... | 60647    | 14987   | nan | MR. BROWN'S ... | 2301 W CHICA... |
| 1    | Proco "Joe" ... | 2058 NORTH W... | 60647    | 15642   | 814 | Beat Kitchen    | 2000–2100 W ... |

Dokleić rzeczy z innej części

`merge_asof()` działa podobnie jak `merge_ordered()` left join.

Dopasowywuje kolumny najbliższej wartości, a nie równe wartości. Wartość ta jest nie większa niz wartość C. Trzeba pamiętać aby kolumny były posortowane. Zastosowanie argumentu `direction=’forward’` spowoduję że wartości przypisane będą więszą lub równe danej wartosci. Natomiast `direction=’nearest’` przypiszę wartośc najbliższą.

![Untitled](Data%20Science%20eefd0a54ac2b4175b68207e2049e3753/Untitled%209.png)

Zastosowanie gdy dane godzinowe się niezgadzają.

```python
pd.merge_asof(visa, ibm, on=['date_time'],
							suffixes=('_visa','_ibm'),
							direction='forward')
```

### Prawdopodobieństwo

Aby wyciągnąć jedną przykładową daną korzystamy z metody `.sample()`. Działa ona losowo

```python
print(sales_counts)

#Output
#    name  n_sales
#0   Amir      178
#1   Brian      128
#2   Claire       75
#3   Damian       69
```

```python
sales_counts.sample()

#Output
#   name  n_sales
#1  Brian      128
```

```python
#    name  n_sales
#2  Claire       75
```

Jeśli użyliśmy metody `sample()` dany rekordu nie ma już we zbiorze.

Jeśli wybiezrszemy to samo ziarno do generowania liczb losowych otrzymamy te same wyniki.

```python
np.random.seed(10)
sales_counts.sample()

#Output
#    name  n_sales
#1  Brian      128
```

Można także podać ilość zwracanych rekordów.
```python
sales_counts.sample(2)
#Output
#     name    n_sales
#1   Brian        128
#2  Claire         75
```


Można przywrócić starą wersję tabeli, korzystając z parametru `replace = True`. 
Jest to losowanie ze zwracaniem.

```python
sales_counts.sample(5, replace = True)
```

Zliczenie wystąpienie każdej wartości w danej tabeli.
```python
    amir_deals["product"].value_counts()
```

Wizualizacja danych na histogramie
```python
rolls_10['number'].hist(bins=np.linspace(1,7,7))
plt.show()
```

### Distributions Continous

Aby obliczyć prawdopodobieństwo korzystamy z metody `.cdf()`
```python
from scipy.stats import uniform
uniform.cdf(7, 0, 12)
#0.5833333
```

`cdf(7, 0, 12)` Oblicza dystrybuantę rozkładu jednostajnego w punkcie 7, dla rozkładu zdefiniowanego od 0 do 12. Oblicza to prawdopodobieństwo że zmienna losowa na przedziale [0, 12] przyjmie wartość mniejszą lub równą 7. 

Jeśli chcemy obliczyć prawdopodobieństwo że liczba będize większa od danej wartości to musimy odjąć od jedności.

```python
from scipy.stats import uniform
1 - uniform.cdf(7, 0, 12)
```

Aby obliczyć przwdopodobieństwo wylosowania między dwoma przedziałami korzystamy z odjęcia większej liczby od mniejszej.

```python
from scipy.stats import uniform
uniform.cdf(7, 0, 12) - uniform.cdf(4, 0, 12)
```

Jeśli chcemy wylosować 10 liczb z przedziału od 0 do 5. 

```python
from scipy.stats import uniform
uniform.rvs(0, 5, size=10)
#Output
#array([1.89740094, 4.70673196, 0.33224683, 1.0137103 , 2.31641255,       3.49969897, 0.29688598, 0.92057234, 4.71086658, 1.56815855])

```

Możliwe jest zasumulowanie w pythonie rzuceniem monetą.

```python
binom.rvs(1, 0.5, size=8)

#Output
#array([0, 1, 1, 0, 1, 0, 1, 1])
```

Powyższa funkcja symuluje rzut monetą gdzie możlie są dwie możliwości albo będzie 0 albo 1.


```python
from scipy.stats import binom
binom.rvs(3, 0.5, size=10)

#Output
#array([0, 3, 2, 1, 3, 0, 2, 2, 0, 0])
```

Powyższy kod mówi że chce rzucić **10 razy** **trzema** monetami z prawdopodobieństwem sukcesu **50%**.

Jeśli chcemy zmienić pradwopodobieństwo sukcesu(w poniższym przypadku je zmniejszyć do 25%) 

```
binom.rvs(3, 0.25, size=10)
```

Prawdopodobieńswo rzucenia 7 orłów na 10 monet można użyć metody `.pmf(7, 10, 0.5)`
```python
binom.pmf(7, 10, 0.5)

#Output
#0.1171875
```

Aby poznać pradopodobieńswo wyrzucenia 7 lub mniej orłów. 

```python
binom.cdf(7, 10, 0.5)

#Output
#0.9453125
```

Prawdopodobieństwo wyrzucenia 7 lub więcej. Wyliczamy za pomocą poniższego wyrażenia.

```python
1 - binom.cdf(7, 10, 0.5)

#Output
#0.0546875
```
   
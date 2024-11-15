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

.pmf() używana jest do obliczenia prawdopodobieństwa wystąpienia konkretnej liczby sukcesów w serii prób Bernulliego

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

Prawdopodobieństwo wyrzucenia więcej niż 7. Wyliczamy za pomocą poniższego wyrażenia.

```python
1 - binom.cdf(7, 10, 0.5)

#Output
#0.0546875
```
   

### Normalna dystrybuanta

Właściwości
- Jest symetryczna
- powierzchnia równa 1
- nigdy nie dotyka zera
- służy do określenia średniej i/lub odchylenia standardowego
- W rozkładzie normalnym 68% danych znajduje sie pomiędzy odchyleniami standardowymi
- 95% wszystkich danych znajduję się w odległości 2 odchyleń standardowych od środka
- 99.7% wszystkich danych znajduje się w odległości 3 odchyleń standardowych od średniej


**Przykład:**
Wyświetlenie informacji ile kobiet ma wzrost mniejszy niż oblicza się:

```python
from scipy.stats import norm
norm.cdf(154, 161, 7)
```

Pierwszą liczbą jest żądana liczba którą chcemy wiedzieć ile osób ma mniejszy wzrost niż 154 cm, drugi parametr to średni wzrost, trzeci parametr to wartość odchylenia standardowego.

Żeby obliczyć ile osób ma wzrost powyżej 154 cm oblicza się za pomocą odjęcia od jedności:

```python
from scipy.stats import norm
1 - norm.cdf(154, 161, 7)
```

Żeby obliczyć ile procent kobiet jest w przedziale między 154 a 157 wyznacza się po przez odjęcie od większego zbioru mniejszego czyli odjęcie 154cm od 157cm.

```python
norm.cdf(157, 161, 7) - norm.cdf(154, 161, 7)
```

Aby wyznaczyć wzrost mniejszy niż u 90% kobiet. Wyznacza się za pomocą:

```python
norm.ppf(0.9, 161, 7)
```

Aby wyznaczyć wzrost większy dla 90% kobiet. Wyznacza się za pomocą:

```python
norm.ppf((1-0.9), 161, 7)
```

Wygenerowanie losowych 10 wzrostów wyznacza się za pomocą poniższego kodu:
```python
norm.rvs(161, 7, size=10)
```

Gdzie pierwszy parametr to średni wzrost, drugi parametr o odchylenie standardowe, a trzeci to liczba generowanych wartości.

**Centralne twierdzenie graniczne** *(ang. Central limit theorem)*

Rozkład próbkowania będzie się zbliżał do rozkładu normalnego w miarę wzrostu liczby prób. Trzeba pamiętać że ma to zastosowanie wtedy i tylko wtedy gdy próbki są podejmowane losowo i są niezależne.

```python
sales_team = pd.Series(["Amir", "Brian", "Claire", "Damian"])
sales_team.sample(10, replace=True)

#array(['Claire', 'Damian', 'Brian', 'Damian', 'Damian', 'Amir', 'Amir', 'Amir',       'Amir', 'Damian'], dtype=object)

#10% wyników to Claire

sales_team.sample(10, replace=True)

#array(['Brian', 'Amir', 'Brian', 'Claire', 'Brian', 'Damian', 'Claire', 'Brian',       'Claire', 'Claire'], dtype=object)

#40% wyników to Claire
```

![alt text](image.png)
Wartość prawdopodobienstwa wylosowania *Claire* wynosi 25%. Ma to przełożenie na wykres.
 

**Rozkład Poissona**

Proces w którym zdarzenia zachodzą z określoną częstotliwością, ale całkowicie losowo.
Rozkład Poissona opisuje prawdopodobieństwo wystąpienia pewnej liczby zdarzeń w ustalonym okresie.

Na przykład: 
Liczba zwierząt adoptowanych co tydzień ze schroniska  

Rozkład Poissona opisuje liczba *lambda*

Liczba lambda opisuje średnią liczbę zdarzen w danym okresie. Wartość ta jest jednocześnie wartością oczekiwaną *(ang. value of the distribution)*

**Rozkład Poissona**


Prawdopodobieństwo, że liczba adopcji będzie równa 5 przy średniej tygodniowej średniej liczbie adopcji wynoszącej 8. 

```python
from scipy.stats import poisson
poisson.pmf(5, 8)
#Output
#0.09160366
```

Aby obliczyc prawdopodobienstwo że liczba adopcji będzie większa niż 5 przy średniej wynoszącej 8. Oblicza się za pomocą ponizszego wzoru:

```python
1 - poisson.cdf(5, 8)
```
Gdy średnia liczba adopcji ulegnie zwiększeniu to prawdopodobinstwo adopcji w takiej samej wartosci co poprzednio ulegnie zwiększeniu. 


Można także wylosować określoną liczbę wartości uwzględniając **rozkład Poissona** 

```python
from scipy.stats import poisson
poisson.rvs(8, size=10)
```

Powyższy przykład przedstawia przypadek gdy wybieramy 10 przykladowych rekordów z rozkładu Poissona gdzie średnia wynosi 8.  


**Rozkład wykładniczy**

Reprezentuje prawdopodobienstwo upływu określonego czasu pomiędzy zdarzeniami Poissona. 

Na przkład: Prawdopodobieństwo uplywu więcej niż jednego dnia pomiedzy adopcjami. Rozkład wykładniczy wykorzystuję wartosc lambda ktora określa szybkość. W przeciwieństwie o rozkładu Poissona jest ciągły, ponieważ reprezentuje czas. 

Oczekiwany czas miedzy żadaniami wyznacza się jako 1/(lambda). Umieszcza się go w wyrażeniu który służy do jego liczenia.


Czas oczekiwania mniejszy niż minuta wyznacza sie po przez wykorzystanie ponizszego wzoru:
```python
from scipy.stats import expon
expon.cdf(1, scale = 2)
```
Gdzie:
Pierwszy argument to wartość dla ktorej chcemy policzyć prawdopodobienstwo a druga to czas między żądaniami. 

Prawdopodobieństwo oczekiwania wynoszace więcej niz 4 minuty wyznacza się za pomocą:
```python
from scipy.stats import expon
1 - expon.cdf(4, scale = 2)
```

Prawdopodobienstwo wynoszące między jedną minuta a czterema minutami wyznacza się za pomocą:

```python
from scipy.stats import expon
expon.cdf(4, scale = 2) - expon.cdf(1, scale = 2)
```

**Rozkład t-Studenta**
Rozkład t-Studenta charakteryzuje się obserwacjami które są oddalone od średniej. Rozkład t ma parametr, który wpływa na grubość ogonów rozkładu. Niższe stopnie swobody powodują grubsze ogony i większe odchylenie standardowe. Wraz ze wzrostem liczby stopni swobody rozkład bardziej przypomina rozkład normalny.

**Rozkład logaryticzno-normalny**

Zmienne które posiadają rozklad logarytmiczno-normalny mają logarytm o rozkładzie normalnym. 

### Korelacja (Miara Pearsona)

Korelacja to badanie zależności między dwiema zmiennymi. Je wartość jest leży pomiędzy-1 do 1. Bada siłę relacji. Minus lub plus oznaczają kierunek zależności. 

Dodatnia wartość korelacji oznacza że jak  wartości osi x rosną to wartości na drugiej osi y też rosną.

Ujemna wartość korelacji oznacza jak wartości na osi x rosną to wartości na osi y maleją.

Wykres przykładowej korelacji:

//TO DO  Wyjaśnić jak to zaszło co to jest **lmplot**

```python
import seaborn as sns
sns.lmplot(x="sleep_total", y="sleep_rem", data=msleep, ci=None)
plt.show()
```

Wyznaczenie współczynnika korelacji

```python
msleep['sleep_total'].corr(msleep['sleep_rem'])
#Output
#0.751755
```

Wartości współczynnika korelacji się nie znmieniają.
```python
msleep['sleep_rem'].corr(msleep['sleep_total'])
#Output
#0.751755
```

**Logarytmiczna korelacja**

//TO DO dodać screeny z 1:38 z filmiku Correlation caveats.

Liniowa korelacja daje słaby wynik. PAMIĘTAJ zawsze twórz wykres by lepiej można było zobaczyć.

```python
msleep['bodywt'].corr(msleep['awake'])

#Output
#0.31
```



```python
msleep['log_bodywt'] = np.log(msleep['bodywt'])

sns.lmplot(x='log_bodywt',
            y='awake',
            data=msleep,
            ci=None)

plt.show()

msleep['log_bodywt'].corr(msleep['awake'])

#Output
#0.57
```

Inne przykłady:
- korelacja kwadratowa,
- korelacja odwrotnie proporcjonalna **1/x**.

Korelacja nie zawiera związku przyczynowo-skutkowego.
Oznacza to że wartość osi x nie powodują zmian wartości na osi y.


## Seaborn
Stworzony by ułatwić tworzenie najpopularniejszych wykresów, działa wyjątkowo dobrze ze strukturami danych, zbudowany na bazie Matplotlib. Jednakże Matplotlib nie jest elastyczny.

Alias `sns` skrót wzięto od **S**amuel **N**orman **S**eaborn. Aktora "The West Wing"

Jeśli chcemy wykorzystać Seaborn trzeba importować również matplotlib gdyż na nim został zbudowany. 

```python
import seaborn as sns
import matplotlib.pyplot as plt

height = [62, 64, 69, 75, 66, 68, 65, 71, 76, 73]

weight = [120, 136, 148, 175, 137,165, 154, 172, 200, 187]

sns.scatterplot(x=height, y=weight)
plt.show()
```

```python
import seaborn as sns
import matplotlib.pyplot as plt

gender = ["Female", "Female", "Female", "Female", "Male", "Male", "Male", "Male", "Male", "Male"]

sns.countplot(x=gender)
plt.show()

```

Należy pamiętać aby DataFrame był uporządkowany. Nie uporządkowana ramka danych nie będzie działać prawidłowo. 


```python
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

df = pd.read_csv("masculinity.csv")
sns.countplot(x="how_masculine",
               data=df)

plt.show()

```

Można rozróniać kolory punktów ze względu na daną kategorię.

Przed dodaniem kodu który odpowiada za kolor

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.scatterplot(x="total_bill",
                y="tip",
                data=tips)

plt.show()
```
Po dodaniu parametru `hue` który odpowiada za kolor. Automatycznie dodaje legendę do naszego wykresu. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.scatterplot(x="total_bill",
                y="tip",
                data=tips,
                hue="smoker")

plt.show()
```

Po dodaniu parametru `hue_order`, która pobiera listę wartości i odpowiednio ustawia kolejność wartości na wykresie.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.scatterplot(x="total_bill",
                y="tip",
                data=tips,
                hue="smoker",
                hue_order=["Yes",
                            "No"])

plt.show()
```

Można do używania kolorów wykorzystywać palety kolorów.

```python
import matplotlib.pyplot as plt
import seaborn as sns

hue_colors = {"Yes" : "black",
               "No" : "red"}

sns.scatterplot(x="total_bill",
                y="tip",
                data=tips,
                hue="smoker",
                palette=hue_colors)

plt.show()
```

Można także używać skrótów a także kodów HTML

```python
import matplotlib.pyplot as plt
import seaborn as sns

hue_colors = {"Yes" : "#808080",
               "No" : "#00FF00"}

sns.scatterplot(x="total_bill",
                y="tip",
                data=tips,
                hue="smoker",
                palette=hue_colors)

plt.show()
```

Można wykorzystywać odcień do rozróżnienia kolorów.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.countplot(x="smoker",
                data=tips,
                hue="sex")

plt.show()
```

`relplot()` umożliwia korzystanie z wykresów punktowych jak i liniowego. Pozwala na tworzenie wielu wykresów na przy użyciu jednego schematu.

Różnice między satterplot() i relplot():

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.scatterplot(x='total_bill',
                y='tip',
                data=tips)

plt.show()
```

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
                y='tip',
                data=tips,
                kind='scatter')

plt.show()
```

Podsumowywując różnice między to to że w relplot wybieramy sobie rodzaj jaki to ma być wykres. Do wyboru typu wykresu służy `kind`.

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
                y='tip',
                data=tips,
                kind='scatter',
                col='smoker')

plt.show()
```

Parametr `col` **NIE** oznacza **koloru** lecz przedstawienie wykresów jeden obok drugiego(kolumnowo) danych które znajdują się w kolumnie `smoker`. Aby przedstawić wykresy jeden pod drugim wykorzystuje się parametr `row`.

Można także wykorzystywać oba te parametry. 

Istnieje możliwość przedstawienia wykresów z wykorzystaniem zawijania *(ang. wrap)*

Zawijanie kolumn

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
                y='tip',
                data=tips,
                kind='scatter',
                col='day',
                col_wrap=2)

plt.show()
```


Można także zastosować uporządkowanie wykresów z wykorzystaniem parametru `col_order`. 

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
                y='tip',
                data=tips,
                kind='scatter',
                col='day',
                col_wrap=2,
                col_order=["Thur",
                           "Fri",
                           "Sat",
                           "Sun"])

plt.show()
```
Można także wyświetlić z podziałem na kolumny i wiersze

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
            y='tip',
            data=tips,
            kind='scatter',
            col='smoker',
            row='time')
plt.show()
```

**Rozmiar punktu**
Do lepszej wizualizacji danych lepiej jest wykorzystywać parametr `size`. 

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
            y='tip',
            data=tips,
            kind='scatter',
            size='size')
plt.show()
```

Można także wykorzystać parametr `hue` który rozróżnia kolorami dane punkty. 

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
            y='tip',
            data=tips,
            kind='scatter',
            size='size',
            hue='size')
plt.show()
```

Innym przykładem jest wykorzystanie prametru `style` który przedstawia na wykresie różne rodzaje punktów.

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
            y='tip',
            data=tips,
            kind='scatter',
            hue='smoker',
            style='smoker')
plt.show()
```

Ustawienie przezroczystości punktów. To zastosowanie jest polecane gdy na wykresie jest wiele nakładających się na siebie punktów.

```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.relplot(x='total_bill',
            y='tip',
            data=tips,
            kind='scatter',
            alpha=0.4)
plt.show()
```

### Liniowe Wykresy relacyjne


W bibliotece seaborna występują dwa typy wykresów relacyjnych:

- wykresy punktowe 
  Wykres ten zakłada że każdy punkt jest niezależną obserwacją
- wykresy liniowe
  Opiera sie na śledzeniu tego samego w czasie
  

Wykorzystanie wykresu punktowego do wizualizacji średniej godzinnej wartości występowania dwutlenku azotu nie jest odpowiednie.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df_mean,
            kind='scatter')
plt.show()
```


Wynika to z faktu że chcemy zwizualizować zmienność w czasie. Dlatego lepszym wyborem jest wykorzystanie wykresu liniowego.


```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df_mean,
            kind='line')
plt.show()
```

Wykorzystując dane, które uwzględniają dodatkowo miejsce lokalizacji pomiaru (wschód, zachód, północ, południe) w danym mieście. 

Przedstawiamy wykres który tworzy różny kolor liń i ich kształtów dla danego regionu.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df_loc_mean,
            kind='line',
            style='location',
            hue='location')
plt.show()
```
Dodanie parametru `markers` i ustawienie go na True Spowoduje wyświetlenie znacznika punktu dla każdego punktu danych. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df_loc_mean,
            kind='line',
            style='location',
            hue='location',
            markers=True)
plt.show()
```

Jeśli nie chcemy aby style lin różniły się od podgrupy to wtedy należy ustawić parametr `dashes` na `False`

Do wizualizacji danych, których jest wiele na każdą godzinę zastosowanie wykresu punktowego `scatter` prowadzi do stworzenia wykresu który jest nie czytelny a punkty zachodzą na siebie. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df,
            kind='scatter'
            )
plt.show()
```

W takich wypadkach należy zastosować wykres liniowy.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df,
            kind='line'
            )
plt.show()
```

Wykres liniowy zawiera wiele obserwacji na każdą wartość **x**, które zostaną zagregowane w jedną sumarycznę miarę. Domyślnie wyświetlana jest średnia. Biblioteka seaborn automatycznie oblicza przedział ufności dla każdej średniej. Jest nim zacieniowany obszar. Zacieniowany obszar zawiera 95% danych. 

Można rozszerzyć ten obszar wykorzystując parametr `ci` i ustawiając go na `sd`. W tedy na wykresie zostaną przedstawione wszystkie dane. Przedział ufności zostanie zwiększony.


```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df,
            kind='line',
            ci='sd')
plt.show()
```

Można go nie umieszczać na wykresie. Wtedy do parametru `ci` przypisujemy wartość `None`.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.relplot(x='hour',
            y='NO_2_mean',
            data=air_df,
            kind='line',
            ci='None')
plt.show()
```
### Wykresy kategoryczne



Zalicza się do nich wykresy zliczeniowe i słupkowe. Służą do porównania między różnymi grupami. Wykres zliczający wyświetla liczbę obserwacji w każdej kategorii. 

#### Wykresy zliczeniowe

Do tworzenia różych typów wykresów kategorycznych wykorzystuje się funkcję `catplot`. Podobnie jak `relplot()` wykorzystuje parametry `col=` oraz `row=`

Przykład zastosowania funkcji `countplot()`

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.countplot(x="how_masculine",
              data=masculinity_data)
            
plt.show()            
```

Zamiast wykorzystywać funkcję `countplot()` wykorzystamy funkcję `catplot()`. Różnią się parametrem `kind`

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="how_masculine",
              data=masculinity_data,
              kind="count")
            
plt.show()            
```

Można także uporządkować wykresy według własnego uznania. Robi się to przy pomocy tablicy zawierającej dane rekordy oraz dodaniem do funkcji 

```python
import matplotlib.pyplot as plt
import seaborn as sns

category_order = ["No answer",
                  "Not at all",
                  "Not very",
                  "Somewhat",
                  "Very"]

sns.catplot(x="how_masculine",
              data=masculinity_data,
              kind="count",
              order=category_order)
            
plt.show()            
```

#### Wykresy słupkowe

Wykresy słupkowe pokazują średnią zmniennej ilościowej dla każdej chcianej kategorii.

Podobnie jak w przypadku wykresów zliczeniowych wykorzystuję się funkcję `catplot()`.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="day",
            y="total_bill",
            data=tips,
            kind="bar")

plt.show()
```

Seaborn automatycznie pokazuje przedziały ufności dla tych średnich. Zawierają one 95% danych. Przedziały te przekazują nam informację o poziomie niepewności jaki jest do tych szacunków.  

Można nie pokazywać tego przedziału ufności. Wykorzystujemy do tego parametr `ci` które ustawiamy na `None`

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="day",
            y="total_bill",
            data=tips,
            kind="bar",
            ci=None)

plt.show()
```

Można zmienić orientację. Ustawiamy wtedy zmienna kategoryczną na zmienną `y`. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(y="total_bill",            
            x="day",
            data=tips,
            kind="bar",
            ci=None)

plt.show()
```

W praktyce zmienną kategoryczną ( w tym przypadku `day`) umieszczamy na osi `x`

### Wykresy pudełkowe

Przedstawia rozkład danych ilościowych. Pudełko przedstawia dane znajdujące się pomiędzy pierwszym a trzecim kwartylem. Linia po środku przedstawia medianę. Służą do porównywania rozkładu zmniennej `ilościowej` do zmniennej `kategorycznej`.

**Tworzenie wykresu pudełkowego**

```python
import matplotlib.pyplot as plt
import seaborn as sns

g = sns.catplot(x="time",
                y="total_bill",
                data=tips,
                kind="box")

plt.show()
```

Można także wykorzystywać `boxplot()`. Praktyczniej lepiej jest korzystać z funkcji `catplot` ze względu na możliwość wykorzystywania parametrów `col`, `row`.

Można także zmieniać kolejność `boxplot`

```python
import matplotlib.pyplot as plt
import seaborn as sns

g = sns.catplot(x="time",
                y="total_bill",
                data=tips,
                kind="box",
                order=["Dinner",
                       "Lunch"])

plt.show()
```

Wykresy pudełkowe mogą zawierać punkty odstające *(ang. outliers)*

Można je pominąć i wogóle ich nie brać pod uwagę. Można to osiągnąć po przez zastosowanie parametru `sym` i przypisaniu do niego wartości `""`.

```python
import matplotlib.pyplot as plt
import seaborn as sns

g = sns.catplot(x="time",
                y="total_bill",
                data=tips,
                kind="box",
                sym="")

plt.show()
```

**Rozstęp międzykwartylowy** **(IQR)** to miejsce pomiędzy pierwszym a trzecim kwartylem.

Punkty odstające pojawiają sie po przekroczeniu 1,5 wartości rozstępu międzykwartylowego. Istnieje możliwość zwiększenia "wąsów" przy wykorzystaniu parametru **whis**. Można modyfikować długość wąsów po przez rozszerzenie przedziału otrzymujemy to po przez **2.0 * IQR** **(whis=2.0)**. Można to także zrobić modyfikując przedział **whis=[5, 96]**.

```python
import matplotlib.pyplot as plt
import seaborn as sns

g = sns.catplot(x="time",
                y="total_bill",
                data=tips,
                kind="box",
                whis=[0, 100])

plt.show()
```

### Point plots

Przedstawiają średnią oraz przedział ufności wynoszący 95%.
Jest podobny do wykresu liniowego lecz różni się tym że punktowy wykres jest zależny od zmniennej **x**, która jest zmienną **kategoryczną**. Wykres liniowy jest zależny o wartości x które są zmiennymi **ilościowymi**.

**Wykres punktowy pozwala na czytelniejsze porównanie wysokości podgrup na jednym wykresie w porównaniu do  boxplot**.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="age",
            y="masculinity_important",
            data=masculinity_data,
            hue="feel_masculine",
            kind="point")

plt.show()
```

Można storzyć wykres, gdzie punkty oznaczające średnie nie mają połączeń między kolejnumi punktami. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="age",
            y="masculinity_important",
            data=masculinity_data,
            hue="feel_masculine",
            kind="point",
            join=False)

plt.show()
```

Istnieje możliwość aby punkt na wykresie był medianą. Wykorzystuje sie do tego parametr **estimator** z biblioteki **numpy**.

```python
import matplotlib.pyplot as plt
import seaborn as sns
from numpy import median

sns.catplot(x="age",
            y="masculinity_important",
            data=masculinity_data,
            kind="point",
            estimator=median)

plt.show()
```

**Lepiej jest stosować medianę ze względu na to że jest bardziej odporna na punkty odstające** . 

Można dodać poziome linie na zakończeniu przedziału ufności. Wykorzystuję się do tego parametr **capsize**. 

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="smoker",
            y="total_bill",
            data=tips,
            kind="point",
            capsize=0.2)

plt.show()
```
Isnieje możliwość wyłączenia przedziału ufności.

```python
import matplotlib.pyplot as plt
import seaborn as sns

sns.catplot(x="smoker",
            y="total_bill",
            data=tips,
            kind="point",
            ci=None)

plt.show()
```

### Modyfikacja wyglądu wykresu

W **seaborn** występują pięć różnych rodzai stylów tła background których chcemy wyświetlić: *white, dark, whitegrid,darkgrid, ticks*. Domyślnym kolorem tła jest biały. 

**whitegrid** umożliwia dodanie szarych poziomych linii na wykresie.
**ticks** dodaje małe kreski na osiach,
**dark** dodaje szare tlo wykresu
**darkgrid** dodaje szare tło na wykresie oraz dodaje poziome linie na wykresie, 

Wykorzystuje się do tego metodę **sns.set_style()**

Możliwe jest wykorzystanie **sns.set_palette()** który zmienia kolory,

//TO DO Dodać screena z 2:49 (screen z kolorami) z filmiku Changing plot style and color

**Występują także palety sekwencyjne**, które są domyślne.

Istnieje możliwość tworzenia własnych palet kolorów. Robi się to po przez stworzenie tablicy z nazwami kolorów lub podanie szesnastkowych kodów kolorów. Aby użyć je w wykresie nalezy użyc nazej tablicy jako parametru naszej metody.

**Zmiana skali wykresu:**

Wykonuje się to po przez użycie metody **sns.set_context()**
Istnieje możliwość zmiany skali **"paper", "notebook", "talk", "poster"**

### Dodanie tytułu i nazw zmiennych

Istnieje zmiany wizualizacji nazw na osiach.

Biblioteka Seaborna dostarcza dwa różne typy obiektów **FacetGrids**, **AxesSubplots**

Aby sprawdzić z jakim typem obiektu mamy doczynienia należy użyć funkcji **type()**

FacetGrids składa się z jednego lub większej liczby AxesSubplots który obsługuje wątki podrzędne, czyli możliwe jest stworzenie wiele rodzai wykresów.

**FacetGrid - relplot(), catplot()
AxesSubplot - scatterplot(), countplot()**

Aby dodać tytuł do wykresu należy użyć metody **g.fig.suptitle("New Title")**

Można zmienić rozmiar wielkości tytułów. Aby dodać tytuł do wykresu należy użyć metody **fig.suptitle("New Title", y=1.03)**

Ustawienie tytułu dla całego wykresu w **FacetGrid** dodaje się przy użyciu metody **g.fig.suptitle("New Title", y=1.03)**

Ustawienie tytułu dla całego wykresu w **AxesSubplot** **g.set_title("New Title", y=1.03")**

```python

g = sns.catplot(x="Region",
                y="Birthrate",
                data=gdp_data,
                kind="box)

g.fig.suptitle("New Title",
                y=1.03)

```

```python

g = sns.boxplot(x="Region",
                y="Birthrate",
                data=gdp_data)

g.set_title("New Title",
                y=1.03)

```

Dodanie dwóch wykresów koło siebie z podziałem na grupy. Do podziału słuzy parametr **col**.

```python

g = sns.catplot(x="Region",
                y="Birthrate",
                data=gdp_data,
                kind="box",
                col="Group")

```

Aby dodać podtytuł dla każdego mniejszego wykresu wykorzystuje się metodę **set_titles("This is {col_name}")**
 Ponadto jeśli ma byc to nazwa wartośści z kolumny to trzeba wykorzystać **col_name **

```python

g = sns.catplot(x="Region",
                y="Birthrate",
                data=gdp_data,
                kind="box",
                col="Group")

g.fig.suptitle("New Title",
                y=1.03)

g.set_titles("This is {col_name}")
```

Do ustawienia nazw osi wykorzystuję się metodę **set()**

```python

g = sns.catplot(x="Region",
                y="Birthrate",
                data=gdp_data,
                kind="box")

g.set(xlabel="New X label",
      ylabel="New Y label")

plt.show()
```

Aby obrócić nazwy zmiennych kategorycznych wykorzystujemy metodę **xticks** z parametrem **rotation**.

```python

g = sns.catplot(x="Region",
                y="Birthrate",
                data=gdp_data,
                kind="box")

plt.xticks(rotation=90)
plt.show()
```

## Introdution to function in Python

### Definiowanie funkcji bez przyjmowania żadnej wartości oraz bez zwracania żadnej wartości.

```python

def square():
    new_value = 4 ** 2
    print(new_value)

square()

```

### Definiowanie funkcji z przyjmowaniem wartości ale bez zwracania wartości.

```python

def square(value):
    new_value = value ** 2
    print(new_value)

square(4)
```

### Definiowanie funkcji z przyjmowaniem wartości oraz zwracająca wartość. 

```python
def square(value):
    new_value = value ** 2
    return new_value

num = square(4)

print(num)

```

Opisywanie funkcji przed ciałem funkcji należy napisać w komentarzu co ta funkcja robi.

```python

def square(value):
    """Returns the square of a value."""
    new_value = value ** 2
    return new_value
```

Funkcje przyjmujące wiele wartości:

```python

def raise_to_power(value1, value2):
    """Raise value1 to power of value2."""
    new_value = value1 ** value2
    return new_value


result = raise_to_power(2, 3)

print(result)

```

Zwracanie wielu zmiennych z funkcji odbywa się przy użyciu krotki.

```python

even_nums = (2, 4, 6)

a, b, c = even_nums
```

Można uzyskać dostęp do krotki tak jak do listy.

```python

even_nums = (2, 4, 6)

print(even_nums[1])
```

```python
second_num = even_nums[1]

print(second_num)
```


```python
def raise_both(value1, value2):
    """Raise value1 to the power of value2 and vice versa."""

    new_value1 = value1 ** value2
    new_value2 = value2 ** value1

    new_tuple = (new_value1, new_value2)

    return new_tuple

result = raise_both(2, 3)

print(result)
```

### Globalne i lokalne
Występują trzy rodzaje zmiennych:

- **Globalne zmienne** są zdefiniowane w głównym ciele skryptu,
- **Lokalne zmienne** są zdefiniowane wewnątrz funkcji,
- **wbudowane zmienne** jak print, type.

Do wszystkich zmiennych można przypisywać nowe wartości.

```python

def func1():
    num = 3
    print(num)

def func2():
    global num
    double_num = num * 2
    num = 6
    print(double_num)

```
W powyższych funkcjach używamy zmiennej lokalnej z jednej funkcji, która jest używana jako również lokalna w drugiej funkcji jako zmiennea lokalna.

Python w poszukiwaniu wartości zmiennej przeszukuję najpierw funkcję wewnetrzną, następnie funkcja otaczającą(jesli taka jest) a na końcu przeszukuje zmienne globalne.

#### Zagnieżdzanie funkcji jednej w drugą.

Jeśli w jednej funkcji mamy duzo rzeczy, które się zagnieżdzają to warto wtedy wykorzystać funkcję zagnieżdzoną.

```python

def mod2plus5(x1, x2, x3):
    """Return the remainder plus 5 of tree values"""

    new_x1 = x1 % 2 + 5
    new_x2 = x2 % 2 + 5
    new_x3 = x3 % 2 + 5

    return(new_x1, new_x2, new_x3)

```

Powyższą funkcję możemy zastąpić poniższą funkcją:

```python

def mod2plus5(x1, x2, x3):
    """Returns the remainder plus 5 of three values."""

    def inner(x):
        """Returns the remainder plus 5 of value."""
        return x % 2 + 5
    
    return (inner(x1), inner(x2), inner(x3))

print(mod2plus5(1, 2, 3))
```

```python

def raise_val(n):
    """Return the inner function."""

    def inner(x):
    """Raise x to the power of n."""
        raised = x ** n
        return raised
    
    return inner

square = raise_val(2)
cube = raise_val(3)
print(square(2), cube(4))

#Output
#4 64


```

Jeśli chcemy zmienić wartość zmiennej w nie tylko w jednej funkcji wewnętrznej ale także w całym zakresie zarówno w danej funkcji jak i jej otoczeniu poprzedzamy nazwę zmiennej **nonlocal**.

```python

def outer():
    """Prints the value of n."""
    n = 1

    def inner():
        nonlocal n
        n = 2
        print(n)

    inner()
    print(n)


outer()
```

Wartość zmiennej wyszukuję się najpierw we funkcji lokalnej, następnie w funkcji otaczającej, pożniej w zmiennych globalnych a na końcu w zmiennych wbudowanych. Jest to metoda **LEGB** od pierwszych liter:

- **L**ocal scope,
- **E**nclosing functions
- **G**lobal
- **B**uilt-in


### Ustawienie domyślnego parametru w funkcji

Definiując argumenty funkcji można odgórnie przypisać jaką wartość ma mieć domyślnie argument.

Można go zmienić w wywołaniu funkcji.

```python

def power(number, pow=1):
    """Raise number to the power of pow."""
    new_value = number ** pow
    return new_value

power(9, 2)
#81

power(9, 1)
#9

power(9)
#9

```


Można także używać listy agrumentów (*args) jako argument funkcji.

```python

def add_all(*args):
    """Sum all values in *args together."""

    # Initialize sum
    sum_all = 0

    # Accumulate the sum
    for num in args:
        sum_all += num
    
    return sum_all

add_all(1)
#1

add_all(1, 2)
#3

add_all(5, 10, 15, 20)
#50

```


Można także wykorzystywać słowniki jako argumenty funkcji używając do tego argument ****kwargs**
Słowniki to struktura danych która przechowuje wartości jako **klucz-wartość**

```python

def print_all(**kwargs):
    """Print out key-value pairs in **kwargs."""

    #Print out the key-value pairs
    for key, value in kwargs.items():
        print(key + ": " + value)

print_all(name="dumbledore", job="headmaster")

#Output
#job: headmaster
#name: dumbledore

```

### Funkcje Lambda

Jest to szybszy sposób pisania funkcji. Używa się do tego słowa kluczowego lambda. Jest to szybki sposób pisania funkcji.

Konsktrukcja
Zapisujemy naszą funkcję i przypisujemy do niej po słowie kluczowym **lambda** dwie nazwy argumentów, po nich stawiamy dwukropek ( : ). Po nich stwiamy wyrażenie określające co funkcja ma zwrócić.  

```python

raise_to_power = lambda x, y: x ** y

raise_to_power(2, 3)

#Output
#8

```

Nazywamy to wtedy funkcjami anonimowymi
Sprawdzmy funkcję map, która przyjmuje dwa : funkcję oraz sekwencję taka jak lista. 

// TO DO jak to działa

```python

nums = [48, 6, 9, 21, 1]

square_all = map(lambda num: num ** 2, nums)

print(square_all)

#Output 
#<map object at 0x103e065c0>
```
Jest to w rzeczwistości obiekt mapy.



```python
print(list(square_all))

#Output
#[2304, 36, 81, 441, 1]
```

//TO Do 

//Dodac opisy z dwóch ostatnich lekcji z kursu  
Introduction to Functions in Python z rozdziału: 
**Lambda functions and error-handling** z lekcji Filter(), Reduce().


Funkacja może zwracać błąd. Aby wyłapyrać wyjątek należy użyć kombinacji **try & except**.

```python

def sqrt(x):
    """Return the square root of a number."""
    try:
        return x ** 0.5
    except:
        print("x must be an int or float")

```

Można także zrobić tak żeby wyłapywała tylko jeden rodzaj błędu.

```python

def sqrt(x):
    """Return the square root of a number."""
    try:
        return x ** 0.5
    except TypeError:
        print("x must be an int or float")

```

Jeśli nie chcemy aby funkcja działała dla liczb ujemnych nalezy użyć **if** i dodatkowo użyć słowa kluczowego **raise** co spowoduje przypisanie do wyjątku

```python

def sqrt(x):
    """Return the square root of a number."""
    if x < 0:
        raise ValueError('x must be non-negative')
    try:
        return x ** 0.5
    except TypeError:
        print("x must be an int or float")
```

## Introduction to iterators



### Playing with iterators

enumerate() - funkcja która przyjmuje dowolny argument iterowalny i zwraca obiekt wyliczeniowy, który składa się z indeksu oraz elementu. 

```python

avengers = ['hawkeye', 'iron man', 'thor', 'quicksilver']
e = enumerate(avengers)
print(type(e))

#Output
#<class 'enumerate'>
```
Gdy chcemy zwrócić liste elementów w postaci listy krotek [(index_elementu, element), (index_drugiego_elementu, drugi_element)] to zwrócimy listę krotek.

```python
e_list = list(e)
print(e_list)

#Output
#[(0, 'hawkeye'), (1, 'iron man'), (2, 'thor'), (3, 'quicksilver')]

```

Domyślnym indeksem iteracji jest 0, ale można to zmienić poprzez przypisanie do parametru `start` odpowiedniej wartości,

```python
avengers = ['hawkeye', 'iron man', 'thor', 'quicksilver']
for index, value in enumerate(avengers):
    print(index, value)

#Output
# 0 hawkeye
# 1 iron man
# 2 thor
# 3 quicksilver

for index, value in enumerate(avengers, start=10):
    print(index, value)

#Output
# 10 hawkeye
# 11 iron man
# 12 thor
# 13 quicksilver
```

Funkcja zip(), który dwie listy umieszcza je w liście krotek i zwraca listę krotek gdzie każda krotka składa się z po jednym elemencie z pierwszej i drugiej listy. 

```python 
avengers = ['hawkeye', 'iron man', 'thor', 'quicksilver']
names = ['barton', 'stark', 'odinson', 'maximoff']
z = zip(avengers, names)
print(type(z))

#Outout
#<class 'zip'>

z_list = list(z)
print(z_list)

#Output
#[('hawkeye', 'barton'), ('iron man', 'stark'), ('thor', 'odinson'), ('quicksilver', 'maximoff')]
```

Zamiast tworzyć liste możemy stworzyć pętle
```python
avengers = ['hawkeye', 'iron man', 'thor', 'quicksilver']
names = ['barton', 'stark', 'odinson', 'maximoff']
for z1, z2 in zip(avengers, names):
    print(z1, z2)

#Output
# hawkeye barton 
# iron man stark 
# thor odinson 
# quicksilver maximoff 

```

Można także użyć operatora **splat** do wydrukowania wszystkich elementów
```python

names = ['barton', 'stark', 'odinson', 'maximoff']
z = zip(avengers, names):
print(*z)

```

### Using iterators to load large files into memory

Jeśli danych jest bardzo dużo można je ładować fragmentami przy wykorzystaniu funkcji `read_csv` przy użyciu parametru chunksize, którego wartość ustawiamy na tyle fragmentów ile chcemy załadować. Jeśli mamy plik CSV z kolumną "x" zawierającą liczbę i chcemy obliczyć sumę wszystkich liczb w tej kolumnie, Ale plik jest za duży by przechować go w pamięci. 

Wykorzystanie chunksize spowoduje że każdy gragment będzie ramka danych (Data Frame). Do **listy** `result` zapisywane są polejne rekordy. 

```python

import pandas as pd
result = []
for chunk in pd.read_csv('data.csv', chunksize=1000):
    result.append(sum(chunk['x']))
total = sum(result)
print(total)

#Output
# 4252532
```

Nie musimy tworzyć osobnej listy wystarczy że na początku zainicjalizujemy jakąś zmienną na 0.

```python
import pandas as pd
total = 0
for chunk in pd.read_csv('data_csv', chunksize=1000):
    total += sum(chunk['x'])
print(total)

#Output
# 4252532
```

### List comprehensions and generators

W tym rozdziale zajmiemy się zastąpieniem zapełnieniem listy przy użyciu pętli for, gdyż są nieefektywne obliczeniowo. Można to zrobić w jednym wierszu kodu. 

```python
nums = [12, 8, 21, 3, 16]
new_nums = []
for num in nums:
    new_nums.append(num + 1)
print(new_nums)

#Output
#[13, 9, 22, 4, 17]
```

Zamiast tego można skożystać z wyrażeń listowych. 
Składnia: w nawiasach kwadratowych wpisujemy wartosci które chcesz utworzyć inaczej zwane wyrażeniami wyjściowymi, po którym następuje wykonanie pętli for, która się pożniej do niej odwołuje.

```python
nums = [12, 8, 21, 3, 16]
new_nums = [num + 1 for num in nums]
print(new_nums)

#Output
#[13, 9, 22, 4, 17]
```

Można stworzyć wyrażenie listowe przy użyciu obiektu zakresu.

```python
result = [num for num in range(11)]
print(result)

#Output
#[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
```

Można wykorzystywać wyrażeń listowych do zastępowania dwóch pętli.

```python
pairs_1 = []
for num1 in range(0, 2):
    for num2 in range(6, 8):
        pairs_1.append((num1, num2))
print(pairs_1)

#Output
#[(0, 6), (0, 7), (1, 6), (1, 7)]
```

Powyższy kod można zastąpić wyrażeniem listowym

```python
pairs_2 = [(num1, num2) for num1 in range(0, 2) for num2 in range(6, 8)]
print(pairs_2)

#Output
#[(0, 6), (0, 7), (1, 6), (1, 7)]
```

Istnieje możliwość dodania *wyrażenia warunkowego* - **if**.

```python 
[num ** 2 for num in range(10) if num % 2 == 0]

#Output
# [0, 4, 16, 36, 64]
```
Listę można uzależnić od wyrażenia wyjściowego.

```python
[num ** 2 if num % 2 == 0 else 0 for num in range(10)]

#Output
# [0, 0, 4, 0, 16, 0, 36, 0, 64, 0]
```

Możemy także pisać wyrażenia listowe tworzące **słowniki** z elementów iteracyjnych. Różnicą jest to że używamy nawiasów klamrowych zamist kwadratowych.

```python

pos_neg = {num: -num for num in range(9)}
print(pos_neg)

#Output
#{0: 0, 1: -1, 2: -2, 3: -3, 4: -4, 5: -5, 6: -6, 7: -7, 8: -8}

print(type(pos_neg))
#Output
#<class 'dict'>

```

Generatory jest obiektem który nie tworzy listy ale można po nim iterować

```python
result = (num for num in range(6))
for num in result:
    print(num)

#Output
# 0
# 1
# 2
# 3
# 4
# 5
```

Można przekazać generator do listy aby utworzyc listę.

```python 
result = (num for num in range(6))
print(list(result))
```

Można iterować generator jeden po drugim

```python
result = (num for num in range(6))

print(next(result))
# 0

print(next(result))
# 1
```
W tym wypadku wartość jest wyciągana tylko wtedy gdy jest potrzebna. Nazywa się oceną leniwą, której wartość wyrażenia jest opóżniona. Może to bardzo pomóc podczas pracy z bardzo duzymi sekwencjami danych ponieważ umożliwia generowanie elementy sekwencji w locie. 

Na generatorach można wykonywać te same operacje co na listach.

Funkcje generatora, które po wywołaniu tworzą obiekty generatora. Są pisane ze składnią dowolnie innej funkcji zdefiniowanj przez użytkownika lecz zamiast zwracania wartości za pomocą slowa kluczowego return zwracają sekwencje wartości przy użyciu słowa kluczowego yield.

```python
def num_sequence(n):
    """Generate values from 0 to n."""
    i = 0
    while i < n:
        yield i
        i += 1
```


## Exploratory Data Analysis in Python

### Getting to Know a Dataset

Aby wczytać plik używamy metody `read_csv` dostępnej w bibliotece Pandas. Wczytanie pierwszych rekordów (domyślnie 5) odbywa się przez wykorzystanie metody `head()`

```python

books = pd.read_csv("book.csv")
books.head()
```

|                          name |                   author |  rating | year |       genre |
|-------------------------------|--------------------------|---------|------|-------------|
| 10-Day Green Smoothie Cleanse |                 JJ Smith |    4.73 | 2016 | Non Fiction |
|             11/22/63: A Novel |             Stephen King |    4.62 | 2011 |     Fiction |
|             12 Rules for Life |       Jordan B. Peterson |    4.69 | 2018 | Non Fiction |
|        1984 (Signet Classics) |            George Orwell |    4.73 | 2017 |     Fiction |
|          5,000 Awesome Facts  | National Geographic Kids |    4.81 | 2019 |   Childrens |

Żeby wyświetlić informację o danych wykorzystujemy metodę `info()`. Jakiego są typu, ile ich jest, czy są wartości null, jakie są nazwy kolumn i ile jest rekordów w każdej kolumnie

``` python

books.info()

#<class 'pandas.core.frame.DataFrame'>
#RangeIndex: 350 entries, 0 to 349
#Data columns (total 5 columns):
#  #     Column  Non-Null Count     Dtype 
# --     ------  --------------     -----  
#  0       name    350 non-null    object
#  1     author    350 non-null    object 
#  2     rating    350 non-null    float64
#  3       year    350 non-null    int64  
#  4      genre    350 non-null    object 
# dtypes: float64(1), int64(1), object(3)
# memory usage: 13.8+ KB
```

Aby zliczyć ile jest rekordów w danej kolumnie używamy metody `value_counts`.

```python
books.value_counts("genre")

#genre
#Non Fiction    179
#Fiction        131
#Childrens       40
#dtype: int64


```

Aby wyświetlić informacje dotyczące opisu danych liczbowych wykorzystujemy metodę `describe()`.


```python
books.describe()

#           rating        year
#count      350.000000    350.000000
#mean       4.608571      2013.508571
#std        0.226941      3.284711
#min        3.300000      2009.000000
#25%        4.500000      2010.000000
#50%        4.600000      2013.000000
#75%        4.800000      2016.000000
#max        4.900000      2019.000000

```

Do wizualizacji danych można wykorzytać histogram.

```python
import seaborn as sns
import matplotlib.pyplot as plt
sns.histplot(data=books, x='rating')
plt.show()
```

Aby wyświetlić dane w nieco uporządkowany sposób może posłużyć parametr `binwidth`. Porządkuje on dane tak by wyświetlały się co zadany przedział.

```python
sns.histplot(data=books, x='rating', binwidth=.1)
```


Aby sprawdzić jakiego typu są dane w kolumnach używamy metody `info()` lub `dtypes`.


Aby zmienić wartości w kolumnie wykorzystujemy metodę `astype()` Dostępne są następujące modywfikacje:
- String `str`
- Integer `int`
- Float `float`
- Dictionary `dict`
- List `list`
- Boolean `bool`


```python
books["year"] = books["year"].astype(int)
books.dtypes
```

Aby sprawdzić czy znajdują się wybrane dane kategoryczne wykorzystujemy metodę `isin()`. Zwraca ona wartość logiczną. Jeśli dana kategoryczna znajduję się w danej kolumnie to zwraca `True` jeśli jej nie ma to zwraca `False`.

```python

books["genre"].isin(["Fiction", "Non Fiction"])
```

Można wyświetlić negację poprzez poprzedzenie powyższego wyrażenia tyldą `~`.

```python

books["genre"].isin(["Fiction", "Non Fiction"])
```

Wyświetlenie tylko tych rekordów które spełniają założenie.

```python
books[books["genre"].isin(["Fiction", "Non Fiction"])].head()
```

Wyświetlenie tylko tych rekordów któey mają ten sam typ co dana kolumna.

```python

books.select_dtypes("number").head()
```

Wyświetlenie wartości minimalnej przy użyciu metody `min()`

```python
books["year"].min()

# Output
# 2009
```

Wyświetlenie wartości maksymalnej przy użyciu metody `max()`

```python
books["year"].max()

# Output
# 2019
```

Dane można grupować po zmiennych kategorycznych występujących w danej kolumnie przy użyciu metody `groupby()`.

```python
books.groupby("genre").mean()
```


Agregacja danych
Wykorzystuje się do tego metodę `agg()` w której umieszczamy funkcje gregujące takie jak:

- sum `sum()`
- count `count()`
- minimum `min()`
- maximum `max()`
- variance `var()`
- Standard deviation `std()`


```python
books.agg(["mean", "std"])
```

Można agregować wybrane kolumny przy użyciu jednych funkcji agregujących oraz inne kolumny przy wykorzystaniu innych funkcji agregujących.

```python
books.agg({"rating": ["mean", "std"], "year": ["median"]})
```

Można jednocześnie grupować oraz agregować dane.

```python
books.groupby("genre").agg(
    mean_rating=("rating", "mean"),
    std_rating=("rating", "std"),
    median_year=("year", "median")
)
```


Można przedstawić wizualizacje danych kategorycznych przy wykorzystaniu wykresu kolumnowego:

```python
sns.barplot(data=books, x='genre', y="rating")
plt.show()
```


### Addressing missing data

Aby zsumować ilość brakujących danych używamy metodę `sum()`

```python
print(salaries.isna().sum())
```

Aby usunąć brakujące wiersze wykorzystuje się metodę `dropna`

W poniższym usuwamy dane w kolumnach w których liczba brakujących danych nie przekracza 5%.

```python
threshold = len(salaries) * 0.05
print(threshold)

cols_to_drop = salaries.columns[salaries.isna().sum() <= threshold]
print(col_to_drop)

salaries.dropna(subset=cols_to_drop, inplace=True)
```

Parametr `subset` określe z jakich kolumn mają być usuwane wiersze/
Parametr `inplace=True` powoduje że zmiany zostaną dokonane bezpośrednio na obiekcie `salaries`, bez konieczności przypisywania wyniku do nowej zmiennej. 

```python
cols_with_missing_values = salaries.columns[salaries.isna().sum() > 0]
print(cols_with_missing_values)

for col in cols_with_missing_values[:-1]:
    salaries[col].fillna(salaries[col].mode()[0])
```

Powyższa pętla iteruje przez każdą kolumnę z wyłączeniem ostatniej kolumny. `salaries[col].fillna(salaries[col].mode()[0])`. metoda `fillna` uzupełnia brakujące wartości w kolumnie `col` wartością mody `mode()[0]`. Indeks `[0]` oznacza że to jest pierwsza najczęściej występująca wartość. (najczęściej występująca wartość dla danej kolumny)


Poniższy kod grupuje według wartości w kolumnie `Experience` a następnię kolumna `Salary_USD` brana jest do dalszych analiz i wyliczane jest mediana, a na końcu zapisujemy to do słownika. 

```python
salaries_dict = salaries.groupby("Experience")["Salary_USD"].median().to_dict()
print(salaries_dict)
```

Można także brakujące dane uzupełniać po przez mapowanie przy wykorzystaniu słownika.

```python
salaries["Salary_USD"] = salaries["Salary_USD"].fillna(salaries["Experience"].map(salaries_dict))
```

Metoda `nunique()` Zlicza ile jest unikalnych stanowisk. 

```python
print(salaries["Designation"].nunique())
```

Metoda `str.contains` przeszukuje kolumnę pod kątem określonego ciągu.

```python
salaries["Designation"].str.contains("Scientist")
```

Aby wyszukać dane z wybranej kolumny  które zaczynają się od frazy wykorzystujemy znak `^`.

```python
salaries["Designation"].str.contains("^Data")
```

```python
job_categories = ["Data Science", "Data Analytics",
                  "Data Engineering", "Machine Learning",
                  "Managerial", "Consultant"]

data_science = "Data Scientist|NLP"
data_analyst = "Analyst|Analytics"
data_engineer = "Data Engineer|ETL|Architect|Infrastructure"
ml_engineer = "Machine Learning|ML|Big Data|AI"
manager = "Manager|Head|Director|Lead|Principal|Staff"
consultant = "Consultant|Freelance"

conditions = [
    (salaries["Designation"].str.contains(data_science)),
    (salaries["Designation"].str.contains(data_analyst)),
    (salaries["Designation"].str.contains(data_engineer)),    
    (salaries["Designation"].str.contains(ml_engineer)),    
    (salaries["Designation"].str.contains(manager)),    
    (salaries["Designation"].str.contains(consultant))
]

salaries["Job_Category"] = np.select(conditions,
                                     job_categories,
                                     default="Other")

print(salaries[["Designation", "Job_Category"]].head())

#                     Designation           Job_Category
# 0                Data Scientist           Data Science
# 1    Machine Learning Scientist       Machine Learning
# 2             Big Data Engineer       Data Engineering
# 3          Product Data Analyst         Data Analytics
# 4     Machine Learning Engineer       Machine Learning

sns.countplot(data=salaries, x="Job_Category")
plt.show()

```
Dane które są obiektami można przekonwertować na dowolny typ.
Najpierw trzeba usunąć przecinki jako seperatory tysięcy, następnie float który zmieni typ danych. Stworzenie nowej kolumny do konwersji danych. 

Aby zastąpić jakiś znak jakimś innym wykorzystujemy metodę `str.replace()`
```python
pd.Series.str.replace("znak do usunięcia", "znak który służy do zastąpienia")
```

Na przykład:
Zastępujemy przecinek i niczym go niezastępujemy.

```python
salaries["Salary_In_Rupees"] = salaries["Salary_In_Rupees"].str.replace(",", "")
```
Aby przekonwertować dane używamy metody `astype()`

```python
salaries["Salary_In_Rupees"] = salaries["Salary_In_Rupees"].astype(float)
```
Wiedząc że Rupee to 0,012 dolara amerykańskiego

```python
salaries["Salary_In_Rupees"] = salaries["Salary_In_Rupees"] * 0.012
```

Aby po grupowac dane po wartościach w kolumne Company_Size tak by zawierały średnie zarobki w dolarach używamy poniższego przykładu:

```python
salaries.groupby("Company_Size")["Salary_USD"].mean()
```

Zamiast tworzyć nową tabelę jak w przypadku powyższym lepiej jest stworzyć kolumnę która będzie wyliczała odchylenie standardowe od wynagrodzeń na podstawie doświadczenia.

```python
salaries["std_dev"] = salaries.groupby("Experience")["Salary_USD"].transform(lambda x: x.std())
```

Aby zliczyc ile jest poszczególnych przypadków korzystamy z metody `value_counts()`

```python
print(salaries[["Experience", "std_dev"]].value_counts())
```

Aby stworzyć kolumnę w której będą przechowywane mediany wynagrodzeń na podstawie wielkości firmy. Robimy to przy wykorzystaniu poniższego wyrażenia.

```python
salaries["median_by_comp_size"] = salaries.groupby("Company_Size")["Salary_USD"].transform(lambda x: x.median())
```

Punkty Odstające górne pojawiają się dla wartości wiekszej niż 1,5 odstępu między kwartylnego (**IQR**) ponad trzeci kwartyl.

IQR = 75th - 25th percentile
Upper Outliers > 75th percentile + (1.5 * IQR)

Punkty odstające dolene pojawiają sie dla wartości mniejszej niż 1.5 odstępu między kwartylnego (**IQR**) poniżej pierwszego kwartylu.

IQR = 75th - 25th percentile
Lower Outliers < 25th percentile - (1.5 * IQR)

```python
seventy_fifth = salaries["Salary_USD"].quantile(0.75)
twenty_fifth = salaries["Salary_USD"].quantile(0.25)
salaries_iqr = seventy_fifth - twenty_fifth

print(salaries_iqr)
```

Aby obliczyc skrajne wartości po których przekroczeniu pojawiają się punkty odstające. Korzystamy z powyższych obliczeń.

```python
upper = seventy_fifth + (1.5 * salaries_iqr)
lower = twenty_fifth - (1.5 * salaries_iqr)

print(upper, lower)
```

Aby wyświetlić dane dla których wartości są odstające korzystamy z poniższego kodu:

```python
salaries[(salaries["Salary_USD"] < lower) | (salaries["Salary_USD"] > upper)][["Experience", "Employee_Location", "Salary_USD"]]
```

Możemy odfiltrować punkty odstające

```python
no_outliers = salaries[(salaries["Salary_USD"] > lower) & (salaries["Salary_USD"] < upper>)]

print(no_outliers["Salary_USD"].describe())
```

### Relationships in Data

### Data Miesiące dzień

W większości przypadków podczas importu danych dane dotyczące dat i godzin są typami `object`.

```python
divorce = pd.read_csv("divorce.csv")
divorce.head()

divorce.dtypes
```

Można naprawic tak by odrazu dane były czytane jako daty używając parametru `parse_dates`

```python

divorce = pd.read_csv("divorce.csv", parse_dates=["marriage_date"])
divorce.dtypes
```

Można także zmienić typ danych po zaimportowaniu danych.

```python
divorce["marriage_date"] = pd.to_datetime(divorce["marriage_date"])
divorce.dtypes
```

Jeśli data jest przechowywana w trzech róznych kolumnach takich jak miesiąc dzień rok to mozna to zapisać we wspólnej kolumnie.

```python
divorce["marriage_date"] = pd.to_datetime(divorce[["month", "day", "year"]])
```

Można wyłuskać same miesiące z daty przy wykorzystaniu parametru `dt.month`

```python
divorce["marriage_month"] = divorce["marriage_date"].dt.month
divorce.head()
```

Można stworzyć wykres zależności miedzy długością małżeństwa a miesiącem gdy wzięto ślub.

```python
sns.lineplot(data=divorce, x="marriage_month", y="marriage_duration")
plt.show()
```

Korelacja opisuje kierunek zależności między dwiema zmiennymi oraz jej wartość(siłę)

```python
divorce.corr()
```

**Heatmapa**

Określa szybka wizualizacje interpretacji korelacji. 

//TO DO
Wyjaśnij co znaczy parametr `annot`

```python
sns.heatmap(divorce.corr(), annot=True)
plt.show()
```

Powyższa korelacja jest liniowa.

Możliwe są korelacja super-liniowe (O(n log n)) lub kwadratowe.

Dlatego należy wzbogacać nasze wykresy korelacji wykresami punktowymi. 

Poniżej przedstawiono wykres korelacji zarobków miesięcznych kobiet i mężczyzn w chwili rozwodu

```python
sns.scatterplot(data=divorce, x="income_man", y="income_woman")
plt.show()
```

// TO DO    Dodać wykres punktowy

Po tej wizualizacji widać że związek istnieje, ale nie jest szczególnie silny tak jak to sugerowała nasza mapa popularności.

Aby zobaczyć wykres wszystkich relacji między zmiennymi numerycznymi umozliwia metoda `pairplot`

```python
sns.pairplot(data=divorce)
plt.show()
```

Można także ograniczyć liczbę wyświetlonych relacji, ustawiając parametr `vars`.

```python
sns.pairplot(data=divorce, vars=["income_man", "income_woman", "marriage_duration"])
plt.show()
```

### Związki między zmiennymi kategorycznymi

```python
divorce["education_man"].value_counts()
```

Należy dokonac wizualizacji by zobaczyć powiązania miedzy liczbą małzeństwa a długością małżeństwa.

```python
sns.histplot(data=divorce, x="marriage_duration", binwidth=1)
plt.show()
```

Aby przedstawić związek (liczbę) między długością małżeństwa w odniesieniu do poziomu wykształcenia mężczyzny.

```python
sns.histplot(data=divorce, x="marriage_duration", hue="education_man", binwidth=1)
plt.show()
```

Związek między długością trwania małżeństwa a poziomem wykształcenia nie jest do końca jasny. Wykresy te rozwiązują wykresy Kernel Density Estimate (**KDE**) dostępne w bibliotece Seaborn.

```python
sns.kdeplot(data=divorce, x="marriage_duration", hue="education_man")
plt.show()
```

// TO DO 

W porównaniu do histogramów środowisko KDE uważa się za bardziej zrozumiałe, zwłaszcza gdy pokazuje wiele dystrybucji. Nalezy zauważyć, że lokalizacja szczytowego czasu trwania małżeństwa dla każdego poziomu męskiego jest bardziej widoczna niz na histogramie.

Można wyłączyć algorytm wygładzenia bo to on psuje poprawny odbiór wykresu.

Algorytm ten może powodować że długość niektórych małżeństw może być krótsza od zera. 

Można to zmienić i pominąć algorytm wygładaający po przez zastosowanie parametru `cut`. Parametr `cut` przekazuje informację jak daleko poza minimalne i maksymalne dane powinna sięgać krzywa podczas stosowania wygładzenia.

```python
sns.kdeplot(data=divorce, x="marriage_duration", hue="education_man", cut=0)
plt.show()
```

//TO DO 
DOdac ten wykres

Wykres ten przedstawia tylko te małżeństwa trwające conajmniej jeden rok.

Jeżeli interesuje nas funkcja rozkładu kumulatywnego można ustawić parametr `cumulative`

```python
sns.kdeplot(data=divorce, x="marriage_duration", hue="education_man", cut=0, cumulative=True)
plt.show()
```

Wykres ten przedstawia prawdopodobiństwo że czas trwania małżeństwa będzie krótszy niż dany okres.

Możemy stworzyć przybliżony wiek zawarcia małżeństwa kobiet i mężczyzn.

```python
divorce["man_age_marriage"] = divorce["marriage_year"] - divorce["dob_man"].dt.year
divorce["woman_age_marriage"] = divorce["marriage_year"] - divorce["dob_woman"].dt.year
```

Przedstawiono wykres punktowy na wykresie punktowym.

```python
sns.scatterplot(data=divorce, x="woman_age_marriage", y="man_age_marriage)

plt.show()
```

Przedstawiając dane na wykresie punktowym kolorujące punkty po poziomie wykształcenia przedstawia wykres poniższy:

```python
sns.scatterplot(data=divorce,
                x="woman_age_marriage",
                y="man_age_marriage",
                hue="education_man")

plt.show()
```


Można określić jaka częścią całości jest dana kategorii.
Pozwala to określić czy dany zbiór jest reprezentacyjny dla całej populacji. 
```python
planes["Destination"].value_counts(normalize=True)
```

Inną metodą na zwerfikowanie częstość kombinacji klas są tabele krzyżowa. W naszym przypadku to częstotliwość tras lotów.

Następnie wybieramy kolumnę, która ma być używana jako indeks dla tabeli 
```python

pd.crosstab(planes["Source"], planes["Destination"])
```
Można zastosować dodatkowo funkcje agregujące. W poniższym przypadku jest to mediana cen dla danego kierunku

```python
pd.crosstab(planes["Source"], planes["Destination"],
            values=planes["Price"], aggfunc="median")
```

Jeśli mam dane dotyczące liczbę przystanków wyrażona jako
"1 przystanek"(1 stop), "2 przystanki"(2 stops), "3 przystanki"(3 stops)

```python
print(planes["Total_Stops"].value_counts())

planes["Total_Stops"] = planes["Total_Stops"].str.replace(" stops", "")

planes["Total_Stops"] = planes["Total_Stops"].str.replace(" stop", "")

planes["Total_Stops"] = planes["Total_Stops"].str.replace(" non-stop", "0")

planes["Total_Stops"] = planes["Total_Stops"].astype(int)
```

Aby zobaczyć jak ceny się zmieniają w ciagu miesięcy oraz jak to wygląda w tygodniu.

```python
planes["month"] = planes["Date_of_Journey"].dt.month
planes["weekday"] = planes["Date_of_Journey"].dt.weekday
print(planes[["month", "weekday", "Date_of_Journey"]].head())
```

Można także wyodrębnić godzinę odlotu i przylotu

```python
planes["Dep_Hour"] = planes["Dep_Time"].dt.hour
planes["Arrival_Hour"] = planes["Arrival_Time"].dt.hour
```
Można także grupować dane liczbowe i oznaczać je jako klasy. Czyli stworzenie nowej kolumny która będzie zawierała opis dla danej liczby przypisując jej wartość kategoryczną.

```python
twenty_fifth = planes["Price"].quantile(0.25)
median = planes["Price"].median()
seventy_fifth = planes["Price"].quantile(0.75)
maximum = planes["Price"].max()
```
Następnie tworzymy listę kategorii, a następnie tworzymy zbiór zakresów dla których przypisana zostanie odpowiednia wartość kategorii.

```python
labels = ["Economy", "Premium Economy", "Business Class", "First Class"]
bins = [0, twenty_fifth, median, seventy_fifth, maximum]
```

Do przypisania nowej kategorii z wcześniej utworzonej listy sluży metoda `cut`

```python
planes["Price_Category"] = pd.cut(planes["Price"],
                                  labels=labels,
                                  bins=bins)
```

Możemy teraz stworzyc wykres kategorii lotów w odniesieniu do lin lotniczych.

```python
sns.countplot(data=planes, x="Airline", hue="Price_Category")
plt.show()
```

## Working with Categorical Data in Python

Zmienna jest zmienną kategoryczną o ile zawiera skończoną liczbę odrębnych grup lub kategorii. W badaniach naukowych zwana również danymi jakościowymi. 

Dane kategoryczne można podzielić na trzy typy:
porządkowe i nominalne. 

Jeśli zmnienna kategoryczna ma naturalną kolejność rang naprzykład stopnie wojskowe to można ją uznać za zmieną kategoryczną porządkową. 

Zmienna kategoryczna nominalna to takie dane których nie można umieścić w porządku naturalnym np. ulubione kolory

Jeśli kolumny ramki danych po wykorzystaniu metody `.info()` prezentują dane w kolumnie jako `object` to najprawdopodobniej oznacza że mamy doczynienia ze zmienną kategoryczną. 

Aby otrzymać jeszcze więcej informacji na temat danych w danej kolumnie możemy użyć do tego parametru `describe()` `adult["Marital Status"].describe()`. 
```python
#count                   32561
#unique                      7
#top        Married-civ-spouse
#freq                    14976
#Name: Marital Status, dtype: object
```
Dzięki temu dostajemy informacje o licznie elementów w danej kolumnie a także ile unikalnych wartości znajduje sie w kolumnie, najpopulaniejszą kategorię i liczbę jej wystąpień. 

Innym sposobem jest wykorzystanie metody `value_counts()`, która zwraca wartości kategoryczne oraz ich liczbę.
```python
# Married-civ-spouse       14976
# Never-married            10683
# Divorced                  4443
# Separated                 1025
# Widowed                    993
# Married-spouse-absent      418
# Married-AF-spouse           23
# Name: Marital Status, dtype: int64
```
Można także zwrócić ile procent danej kategorii jest w danej kolumnie przy użyciu metody i odpowiedniego parametru `value_counts(normalize=True)`
```python
# Married-civ-spouse       0.459937
# Never-married            0.328092
# Divorced                 0.136452
# Separated                0.031479
# Widowed                  0.030497
# Married-spouse-absent    0.012837
# Married-AF-spouse        0.000706
# Name: Marital Status, dtype: float64

```

Wyświetlenie typu każdej kolumny odbywa sie po przez wykorzystanie metody `dtypes`

Aby wyświtlic typ danej serii czyli danej kolumny używamy metody `dtype`. `O` oznacza że to obiekty.

```python
adult["Marital Status"].dtype
#dtype('O')

adult["Marital Status"] = adult["Marital Status"].astype("category")

adult["Marital Status"].dtype
#
```

Są dwa sposoby tworzenia serii kategoeycznych, gdy dane nie są jeszcze w formacie DataFrame. Wykorzystuje sie do tego parametr `dtype` 

```python
my_data = ["A", "A", "C", "B", "C", "A"]

my_series1 = pd.Series(my_data, dtype="category")
print(my_series1)
```

Innym sposobem jest wykorzystanie metody `Categorical`
```python
my_data = ["A", "B", "C", "B", "C", "A"]
my_series2 = pd.Categorical(my_data, categories=["C", "B", "A"], ordered=True)
my_series2

#[A, A, C, B, C, A]
#Categories (3, object): [C < B < A]


```

Przechowywanie serii pand z typem categorical jest przydatne z powodu oszczędności ogromnej liczby pamięci.

```python
adult = pd.read_csv("data/adult.csv")
adult["Marital Status"].nbytes
#260488

adult["Marital Status"] = adult["Marital Status"].astype("category")
adult["Marital Status"].nbytes
#32617
```

Można przed wczytaniem zrobic słownik gdzie kluczem będzie kolumna a wartościa typ zmiennej.

```python
adult_dtypes = {"Marital Status": "category"}

adult = pd.read_csv("data/adult.csv", dtype=adult_dtypes)

adult["Marital Status"].dtype

# CategoricalDtype(categories=[' Divorced','Married-AF-spouse',..., ' Widowed'], ordered=False)

```

`.str` umożliwia wykonywanie operacji na danych tekstowych w kolumnach typu `Series`

`.apply()` pozwala na zastosowanie funkcji dla każdego elementu.

Używanie funkcji NumPy w seriach kategorycznych zwykle powoduje błędy. 
Trzeba sprawdzć zawsze czy mamy doczynienia ze zmieną kategoryczną. 

Check - sprawdzenie
```python
used_cars["color"] = used_cars["color"].astype("category")
used_cars["color"] = used_cars["color"].str.upper()
print(used_cars["color"].dtype)
# object
```

convert - przekształcenie

```python
used_cars["color"] = used_cars["color"].astype("category")
print(used_cars["color"].dtype)
# category
```

Trzeba sprawdzać czy zmiany zadziałały zgodnie z zamierzeniami

```python
used_cars["color"] = used_cars["color"].astype("category")
used_cars["color"].cat.set_categories(["black", "silver", "blue"], inplace=True)
used_cars["color"].value_counts(dropna=False)
```

Przykład że nie można wykorzystać funkcji NumPy:

```python
used_cars["number_of_photos"] = used_cars["number_of_photos"].astype("category")
used_cars["number_of_photos"].sum()
```

Można jednak przekonwertować szereg na liczbę całkowitą i zastosować sumowanie.

```python
used_cars["number_of_photos"].astype(int).sum()

# .str converts the column to an array

used_cars["color"].str.contains("red")

# 0     False
# 1     False
```

Kodowania etykiet to technika kodowania wartości kategorycznych jako liczb całkowitych. 
Kody zaczynają się od `0` i kończą na `n - 1`, gdzie n to liczba kategorii.  
Kod `-1` skazywany jest dla brakującej wartości.  Kodowanie etykiet służy do
oszczędzania pamięci i upraszczania odpowiedzi.  

#### Tworzenie kodów

```python
used_cars["manufacturer_name"] = used_cars["manufacturer_name"].astype("category")
used_cars["manufacturer_name"] = used_cars["manufacturer_name"].cat.codes

print(used_cars[['manufacturer_name', 'manufacturer_code']])

#         manufacturer_name     manufacturer_code
# 0                  Subaru                    45
# 38526            Chrysler                     8
```

Przy kodowaniu etykiet należy stworzyć mapy. Można to zrobić tworzyć obiekt dla kodów 
i obiekt dla kategorii.

```python
codes = used_cars["manufacturer_name"].cat.codes
categories = used_cars["manufacturer_name"]

name_map = dict(zip(codes, categories))
```

Używanie słownika:

```python
used_cars["manufacturer_code"] = used_cars["manufacturer_name"].cat.codes
used_cars["manufacturer_code"].map(name_map)
```
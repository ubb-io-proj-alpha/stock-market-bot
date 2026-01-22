# Stock Market Bot
<img src="assets/logo.png" alt="Logo" width="250"/>

Celem projektu jest zbudowanie bota śledzącego akcje na giełdzie i odpowiednie reagowanie na nie, wedle założonej strategii inwestycyjnej.

### Wymagania funkcjonalne
System umożliwia pobieranie danych rynkowych z wielu giełd jednocześnie za pośrednictwem ich interfejsów API oraz ich normalizację do wspólnego formatu. Na podstawie otrzymanych danych system analizuje sytuację rynkową przy użyciu predefiniowanych strategii handlowych i podejmuje decyzje inwestycyjne, takie jak kupno, sprzedaż lub brak akcji. Wygenerowane decyzje handlowe są kolejkowane i następnie przekazywane do realizacji poprzez wysyłanie zleceń do odpowiednich giełd.

Dodatkowo w ramach projektu dostępne są testowe interfejsy API giełdowe wraz z testowym portfelem, na których bot jest uruchamiany i weryfikowany bez użycia rzeczywistych środków finansowych.

### Wymagania niefunkcjonalne
Program wykorzystuje dostępne zasoby systemowe w sposób efektywny, minimalizując opóźnienia w przetwarzaniu danych giełdowych i zapewniając szybką reakcję na zmiany rynkowe.
Zastosowany język programowania Golang pozwala na łatwe wykorzystanie równoległości dzięki goroutines, które są lekkimi wątkami zarządzanymi przez runtime Go. Goroutines korzystają domyślnie z asynchronicznego I/O i są planowane przez scheduler Go na wszystkie dostępne wątki systemowe, co minimalizuje koszty przełączania kontekstu i umożliwia efektywne wykorzystanie procesorów podczas równoległego wykonywania wielu operacji sieciowych I/O oraz jednoczesnego przetwarzania otrzymanych danych. Wąskim gardłem w systemie będą jedynie ograniczenia wynikające z rate-limitingu API giełd oraz przepustowości sieci.

### Stos technologiczny
- [Golang](https://go.dev/)

### Architektura
![Architecture](assets/architecture.png)

### Diagram przypadków użycia
![Use case diagram](assets/use_case.png)


### Uruchamianie i testowanie

Projekt na tym etapie nie zawiera jeszcze kodu źródłowego, ale przygotowana została struktura oraz szablony plików do kompilacji, uruchamiania i testowania. Dzięki temu łatwo będzie rozpocząć implementację oraz zachować porządek w repozytorium.


#### a) Uruchomienie głównego programu

Do kompilacji i uruchamiania bota służą skrypty znajdujące się w katalogu `scripts/`:

- **Windows**:
	1. `scripts/compile.bat` – kompilacja projektu (np. z użyciem `go build`).
	2. `scripts/run.bat` – uruchomienie bota (np. `go run ...` lub uruchomienie binarki).
	Przykład użycia w PowerShell:
	```powershell
	.\scripts\compile.bat
	.\scripts\run.bat
	```

- **Unix (Linux / macOS)**:
	1. `scripts/compile.sh` – kompilacja projektu.
	2. `scripts/run.sh` – uruchomienie bota.
	Przykład użycia:
	```bash
	./scripts/compile.sh
	./scripts/run.sh
	```

#### b) Uruchomienie testów

Analogicznie, do uruchamiania testów służą skrypty w katalogu `scripts/`:

- **Windows**:
	- `scripts/test.bat` – uruchamia testy jednostkowe (np. `go test ./...`).

- **Unix (Linux / macOS)**:
	- `scripts/test.sh` – uruchamia testy jednostkowe.

Przykład użycia:
```powershell
.\scripts\test.bat
```
lub
```bash
./scripts/test.sh
```



#### c) Struktura folderów

Repozytorium zostało podzielone na logiczne katalogi, co ułatwi dalszy rozwój projektu:

- `src` – główny kod programu (logika bota, obsługa API giełd, strategie inwestycyjne itp.)
- `tests` – testy jednostkowe i integracyjne, korzystające z mockowanych API
- `mock` – implementacje mocków API giełdowych, wykorzystywane w testach
- `uml` – źródła diagramów UML (architektura, przypadki użycia)
- `assets` – obrazy i inne zasoby wykorzystywane w dokumentacji (np. logo, diagramy)

Każdy katalog zawiera plik `README.md` lub odpowiednie pliki źródłowe po rozpoczęciu implementacji.


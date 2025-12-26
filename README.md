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

### Uruchamianie
```sh
# TODO...
```

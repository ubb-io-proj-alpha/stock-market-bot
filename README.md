# Stock Market Bot
Celem projektu jest zbudowanie bota śledzącego akcje na giełdzie i odpowiednie reagowanie na nie, wedle założonej strategii inwestycyjnej.

### Wymagania funkcjonalne
System umożliwia pobieranie danych rynkowych z wielu giełd jednocześnie za pośrednictwem ich interfejsów API oraz ich normalizację do wspólnego formatu. Na podstawie otrzymanych danych system analizuje sytuację rynkową przy użyciu predefiniowanych strategii handlowych i podejmuje decyzje inwestycyjne, takie jak kupno, sprzedaż lub brak akcji. Wygenerowane decyzje handlowe są kolejkowane i następnie przekazywane do realizacji poprzez wysyłanie zleceń do odpowiednich giełd.

Dodatkowo w ramach projektu dostępne są testowe interfejsy API giełdowe wraz z testowym portfelem, na których bot jest uruchamiany i weryfikowany bez użycia rzeczywistych środków finansowych.

### Wymagania niefunkcjonalne
Bot wykorzystuje w pełni zasoby dostępne na maszynie, aby zminimalizować opóźnienia w przetwarzaniu danych giełdowych i odpowiednio na nie reagować.
Zastosowany język programowania Golang pozwala na łatwe wykorzystanie równoległości dzięki goroutines, które są lekkimi wątkami zarządzanymi przez runtime Go. Goroutines korzystają domyślnie z asynchronicznego I/O i są planowane przez scheduler Go na wszystkie dostępne wątki systemowe, co minimalizuje koszty przełączania kontekstu i umożliwia efektywne wykorzystanie rdzeni procesora — istotne w projekcie ze stosunkowo wysoką liczbą operacji I/O.

### Stos technologiczny
- [Golang](https://go.dev/)

### Architektura
![Architecture](https://www.plantuml.com/plantuml/png/dPJFRjD048VlynIZN813LCM-SK1jKk1G8VxaW0JYC5qFhcjxnzgCLMp4fS_5DU-5Txkrd50BXLjclczdej-kVBvJSjhM5QWqg8BNeYerPeLZeuLO038g3XTl2zxVPuGN2wGTrkHArjKnN4zWSWmt4vW2FCFbSebNKZTYsMhdIsZ8b9GpBhOgfiHtGSQrw09_0kBNsBWa9Np3QcxOuSKtF_VxGxMQizPJAvcpDbChzTPhcnUMyGmtuZZGXuvdMtMad7VnBpoMgmEsFc1XytkNiSEFBRSSOAX36JR-ODFwxF6pk99TK7mJuDTZR9-u8dFB8RDu5yFlLsCA-18_o7N1haV3jejECdMIfvrqwgJ73aGGjiKa1kUIB6CUxbdk69_tLbnDLT7BtK1UmBZoVqwbfvs2Q0w7njH37gxdalQ_CHiq6CdXy6sNMzhVTtX3Dgk6-K-EIsR79QUC6wzzHXrvqCCJ6CGl5Hks1G7uHY3n1Gpmgctm37UDwpcZssCZ-QUH_ct0ETlCVo7-0000)

### Diagram przypadków użycia
![Use case diagram](https://www.plantuml.com/plantuml/png/VOynJiDW38Nt_nGc4moz0YeiY06fKjMbOd4Jq_v9Vhjo74K9Ogd48NeErcuazo8X1THAT7tly_Dxqy9GhGovkx8D1O9jZfwToHPCO2Lc4Y3nIh5vOUSmDb685jsXjamgzba4MC2Y8FtnbbHnkpTjcE1kzlZC3vwwNI8L3UWWEJVa8QzpWcXk4cVmX9gHmRrO1A_EmNSYJ2QtSBtGCjv86ChWpOappuHV_BhKCeN258tMljsVCaiqRUg0wV7UmLdYYCravKbxAa4cq9hxSJGa_UCCL-EKYqZlQKDMeJQZFMQUYxyb4ag3xpMAwoPjzx11JdAfslUn-2NWEBVRVNvq1sXoYedz-EvS_6l_vgR4IHdoBm00)

### Uruchamianie
```sh
# TODO
```

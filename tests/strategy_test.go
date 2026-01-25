package tests

import (
	"testing"
	"time"
	"fmt"
	"bot/src/app"
)

func TestThresholdStrategy(t *testing.T) {

	// ------------   ARRANGE   --------------

	signalChan := make(chan app.Signal, 50)
	strategy := app.ThresholdStrategy {
		Name: "SP500 Strategy",
		Symbol: "SP500",
		Threshold: 1000,
		MarketDataChan: make(chan app.MarketData, 20),
	}
	currentTime := time.Now()

	// uruchom strategie wewnatrz goroutiny
	// output ze strategi bedzie przekazywany
	// do channelu signalChan
	go strategy.Run(signalChan)
	fmt.Printf("Strategy is running\n")

	// ------------     ACT     --------------

	// przekaz event gieldowy z najnowszymi
	// danymi z gieldy do strategii
	strategy.Broadcast(&app.MarketData {
		SourceMarket: "XTB",
		Symbol: "SP500",
		Price: 1500,
		Timestamp: currentTime,
	})

	// zblokowanie obecnego watku
	// na maks 500 ms i oczekiwanie
	// na przyjscie sygnalu ze strategii
	var receivedSignal *app.Signal
	select {
	case s := <-signalChan:
		receivedSignal = &s
	case <-time.After(500 * time.Millisecond):
		receivedSignal = nil
	}

	// ------------   ASSERT   --------------

	if receivedSignal == nil {
		t.Error("Nie otrzymano sygnalu ze strategii")
	}

	if receivedSignal.Symbol != "SP500" {
		t.Errorf("Otrzymany sygnal powinien miec symbol \"SP500\", otrzymano: %v",
				 receivedSignal.Symbol)
	}

	if receivedSignal.Action != app.ACTION_SELL {
		t.Errorf("Otrzymany sygnal powinien miec akcje ACTION_SELL, otrzymano: %v",
			     receivedSignal.Action)
	}

	if receivedSignal.Price != 1500 {
		t.Errorf("Otrzymany sygnal powinien miec cene 1500, otrzymano: %v",
			     receivedSignal.Price)
	}

	if !receivedSignal.Timestamp.After(currentTime) {
		t.Errorf("Otrzymany sygnal powinien miec znacznik czasu wiekszy niz %v, otrzymano: %v",
				 currentTime, receivedSignal.Timestamp)
	}
}

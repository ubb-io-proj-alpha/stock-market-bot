package main

import "time"

/*
* 	Znormalizowane dane giełdowe
*	pochodzą z fetcherów
*/
type MarketData struct {
	sourceMarket  string
	symbol        string
	price         float64
	timestamp     time.Time
}

type ActionType int
const (
	ACTION_BUY ActionType = iota
	ACTION_SELL
)

/*
* 	Decyzja podjęta przez strategię
*/
type Signal struct {
	strategyName  string
	symbol        string
	action        ActionType
	price         float64
	timestamp     time.Time
}

/*
*	Zlecenie wygenerowane poprzez
*	core/engine na podstawie sygnału
*	pochodzącego ze strategii.
*/
type MarketOrder struct {
	targetMarket  string
	symbol        string
	action        ActionType
	amount        float64
	timestamp     time.Time
}

/*
*	Kontekst aplikacji
*/
type Context struct {
	marketDataChan  chan MarketData
	signalChan      chan Signal
	orderQueueChan  chan MarketOrder
	strategies      []Strategy
}

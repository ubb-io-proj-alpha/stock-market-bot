package app

import "time"

/*
* 	Znormalizowane dane giełdowe
*	pochodzą z fetcherów
*/
type MarketData struct {
	SourceMarket  string
	Symbol        string
	Price         float64
	Timestamp     time.Time
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
	StrategyName  string
	Symbol        string
	Action        ActionType
	Price         float64
	Timestamp     time.Time
}

/*
*	Zlecenie wygenerowane poprzez
*	core/engine na podstawie sygnału
*	pochodzącego ze strategii.
*/
type MarketOrder struct {
	TargetMarket  string
	Symbol        string
	Action        ActionType
	Amount        float64
	Timestamp     time.Time
}

/*
*	Kontekst aplikacji
*/
type Context struct {
	MarketDataChan  chan MarketData
	SignalChan      chan Signal
	OrderQueueChan  chan MarketOrder
	Strategies      []Strategy
}

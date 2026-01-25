package app

import "time"

type Strategy interface {
	Broadcast(in *MarketData)
	Supports(symbol string) bool
	Run(out chan<- Signal)
}

/*
* ===== THRESHOLD STRATEGY =====
*/

type ThresholdStrategy struct {
	Name	   		string
	Symbol	   		string
	Threshold  		float64
	MarketDataChan  chan MarketData
}

func (s *ThresholdStrategy) Run(out chan<- Signal) {
	for data := range s.MarketDataChan {
		signal := Signal {
			StrategyName: s.Name,
			Action: ACTION_BUY,
			Symbol: data.Symbol,
			Price: data.Price,
			Timestamp: time.Now(),
		}
		if data.Price > s.Threshold {
			signal.Action = ACTION_SELL
		}

		out <- signal
	}
}

func (s *ThresholdStrategy) Supports(symbol string) bool {
	return symbol == s.Symbol
}

func (s *ThresholdStrategy) Broadcast(in *MarketData) {
	s.MarketDataChan <- *in
}

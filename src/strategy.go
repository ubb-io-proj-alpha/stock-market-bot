package main

import "time"

type Strategy interface {
	broadcast(in *MarketData)
	supports(symbol string) bool
	process(out chan<- Signal)
}

/*
* ===== THRESHOLD STRATEGY =====
*/

type ThresholdStrategy struct {
	name	   		string
	symbol	   		string
	threshold  		float64
	marketDataChan  chan MarketData
}

func (s *ThresholdStrategy) process(out chan<- Signal) {
	for data := range s.marketDataChan {
		signal := Signal {
			strategyName: s.name,
			action: ACTION_BUY,
			symbol: data.symbol,
			price: data.price,
			timestamp: time.Now(),
		}
		if data.price > s.threshold {
			signal.action = ACTION_SELL
		}

		out <- signal
	}
}

func (s *ThresholdStrategy) supports(symbol string) bool {
	return symbol == s.symbol
}

func (s *ThresholdStrategy) broadcast(in *MarketData) {
	s.marketDataChan <- *in
}

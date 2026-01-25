package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func createContext() *Context {
	context := Context {
		marketDataChan:  make(chan MarketData,   100),
		signalChan:      make(chan Signal,       50),
		orderQueueChan:  make(chan MarketOrder,  100),
		strategies:      []Strategy{},
	}

	return &context
}

func appendStrategy(context *Context, strategy Strategy) {
	context.strategies = append(context.strategies, strategy)
}

func main() {
	context := createContext()

	fmt.Println("Initializing...")

	appendStrategy(context, &ThresholdStrategy { "Zlotowki 1", "PLN", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Zlotowki 2", "PLN", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Zlotowki 3", "PLN", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Zlotowki 4", "PLN", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Zlotowki 5", "PLN", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Zlotowki 6", "PLN", 1000, make(chan MarketData, 20) })

	appendStrategy(context, &ThresholdStrategy { "Dollary 1", "USD", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Dollary 2", "USD", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Dollary 3", "USD", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Dollary 4", "USD", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Dollary 5", "USD", 1000, make(chan MarketData, 20) })
	appendStrategy(context, &ThresholdStrategy { "Dollary 6", "USD", 1000, make(chan MarketData, 20) })

	// spawn single goroutine for each defined strategy
	for _, strategy := range context.strategies {
		go strategy.run(context.signalChan)
	}

	/* broadcast market data to all supported strategies
	*/ 
	go func() {
		for data := range context.marketDataChan {
			for _, strategy := range context.strategies {
				if !strategy.supports(data.symbol) {
					continue
				}

				strategy.broadcast(&data)
			}
		}
	}()

	fmt.Println("OK.")
	fmt.Println("Running...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nExiting...")
}

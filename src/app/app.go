package app

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func createContext() *Context {
	context := Context {
		MarketDataChan:  make(chan MarketData,   100),
		SignalChan:      make(chan Signal,       50),
		OrderQueueChan:  make(chan MarketOrder,  100),
		Strategies:      []Strategy{},
	}

	return &context
}

func appendStrategy(context *Context, strategy Strategy) {
	context.Strategies = append(context.Strategies, strategy)
}

func Run() {
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
	for _, strategy := range context.Strategies {
		go strategy.Run(context.SignalChan)
	}

	/* broadcast market data to all supported strategies
	*/ 
	go func() {
		for data := range context.MarketDataChan {
			for _, strategy := range context.Strategies {
				if !strategy.Supports(data.Symbol) {
					continue
				}

				strategy.Broadcast(&data)
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

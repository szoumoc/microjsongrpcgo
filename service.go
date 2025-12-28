package main

import (
	"context"
	"fmt"
	"time"
)

// Pricefetcher is an interface that can fetch a price
type Pricefetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// pricefetcher implements the pricefetcher interface
type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	fmt.Println("hi just here in service")
	return MockPricefetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPricefetcher(ctx context.Context, ticker string) (float64, error) {
	// mock fetching price from 3rd party service
	time.Sleep(100 * time.Millisecond)

	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("The given ticker is not supported")
	}
	return price, nil
}

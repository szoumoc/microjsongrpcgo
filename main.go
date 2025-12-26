package main

import (
	"context"
	"fmt"
)

// Pricefetcher is an interface that can fetch a price
type Pricefetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// pricefetcher implements the pricefetcher interface
type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPricefetcher(ctx, ticker)

}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPricefetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("The given ticker is not supported")
	}
	return price, nil
}

func main() {
	// This is a placeholder for the main function.
}

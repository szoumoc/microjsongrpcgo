package main

import (
	"context"
	"fmt"
)

type Pricefetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	prices := map[string]float64{
		"BTC": 60000.0,
		"ETH": 3000.0,
	}

	price, ok := prices[ticker]
	if !ok {
		return 0, fmt.Errorf("ticker %s not found", ticker)
	}

	return price, nil
}

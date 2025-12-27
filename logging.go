package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next Pricefetcher
}

func NewLoggingService(next Pricefetcher) Pricefetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Hi im here")
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

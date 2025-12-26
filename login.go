package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggingService struct {
	next Pricefetcher
}

func (s *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"err":   err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

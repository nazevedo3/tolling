package main

import (
	"time"

	"github.com/nazevedo3/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleWare struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleWare{
		next: next,
	}
}

func (m *LogMiddleWare) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.New().WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("AggregateDistance")
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return err
}

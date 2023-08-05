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

func (m *LogMiddleWare) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.New().WithFields(logrus.Fields{
			"took":     time.Since(start),
			"err":      err,
			"obuID":    obuID,
			"amount":   amount,
			"distance": distance,
		}).Info("CalculateInvoice")
	}(time.Now())
	inv, err = m.next.CalculateInvoice(obuID)
	return
}

package main

import "log"

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

const kafkaTopic = "obudata"

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	svc = NewLogMiddleWare(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}

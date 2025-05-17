package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type Telemetry struct {
	CPU       float64 `json:"cpu"`
	Timestamp string  `json:"timestamp"`
}

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "telemetry",
	})
	defer writer.Close()

	for {
		data := Telemetry{
			CPU:       rand.Float64()*100 + 1,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		payload, _ := json.Marshal(data)

		if err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: payload,
		}); err != nil {
			log.Println("Failed to write to Kafka:", err)
		} else {
			log.Println("Telemetry sent:", string(payload))
		}

		time.Sleep(2 * time.Second)
	}
}

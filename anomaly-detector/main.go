package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

type Telemetry struct {
	CPU       float64 `json:"cpu"`
	Timestamp string  `json:"timestamp"`
}

var ctx = context.Background()

func getThreshold(rdb *redis.Client) float64 {
	val, err := rdb.Get(ctx, "cpu_threshold").Float64()
	if err != nil {
		log.Println("Threshold fallback to 80.0%")
		return 80.0
	}
	return val
}

func publishAlert(rdb *redis.Client, msg string) {
	rdb.Publish(ctx, "alerts", msg)
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "alert_stream",
		Values: map[string]interface{}{
			"message": msg,
			"time":    time.Now().Format(time.RFC3339),
		},
	})
}

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "redis:6379"})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "telemetry",
		GroupID: "anomaly-detector-group",
	})
	defer reader.Close()

	alerted := make(map[string]bool)

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka error:", err)
			continue
		}

		var metric Telemetry
		if err := json.Unmarshal(msg.Value, &metric); err != nil {
			log.Println("Bad JSON:", err)
			continue
		}

		threshold := getThreshold(rdb)
		if metric.CPU > threshold {
			alertKey := fmt.Sprintf("%.0f@%s", metric.CPU, metric.Timestamp[:16]) // dedup per minute
			if !alerted[alertKey] {
				alert := fmt.Sprintf("⚠️ High CPU: %.2f%% at %s (threshold: %.2f)", metric.CPU, metric.Timestamp, threshold)
				publishAlert(rdb, alert)
				log.Println(alert)
				alerted[alertKey] = true
			}
		}
	}
}

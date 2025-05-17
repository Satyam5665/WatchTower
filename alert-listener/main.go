package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	sub := rdb.Subscribe(ctx, "alerts")

	fmt.Println("🚨 Listening for alerts...")

	ch := sub.Channel()
	for msg := range ch {
		fmt.Println("🔔 ALERT RECEIVED:", msg.Payload)
	}
}

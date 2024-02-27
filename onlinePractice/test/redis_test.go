package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"onlinePractice/models"
	"testing"
	"time"
)

var ctx = context.Background()

var redisServer = redis.NewClient(&redis.Options{
	Addr:     "192.168.239.100:6379",
	Password: "ranzhou",
	DB:       0, // use default DB
})

func TestRedisSet(t *testing.T) {
	redisServer.Set(ctx, "name", "mmc", time.Second*10)
}

func TestRedisGet(t *testing.T) {
	value, err := redisServer.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value)
}

func TestRedisGetByModel(t *testing.T) {
	value, err := models.RedisServer.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value)
}

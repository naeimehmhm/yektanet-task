package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(host string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", host),
	})

	return &Redis{
		client: client,
	}
}

func (r *Redis) Ping() error {
	c := context.Background()
	pong, err := r.client.Ping(c).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New("invalid ping output")
	}
	return nil
}

package redis

import (
	"github.com/go-redis/redis"
)

type redisClient struct {
	Client *redis.Client
}

func NewRedisClient() *redisClient {
	client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
	})

	return &redisClient{Client: client}
}


func (c *redisClient) Set(key string, value []byte) error {
	err := c.Client.Set(key, value, 0).Err()
	if err != nil {
			return err
	}

	return nil
}

func (c *redisClient) Get(key string) ([]byte, error) {
	val, err := c.Client.Get(key).Bytes()
	if err != nil {
			return val, err
	}

	return val, nil
}

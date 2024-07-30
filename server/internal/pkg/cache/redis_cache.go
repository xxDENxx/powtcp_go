package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/xxDENxx/powtcp_go/server/internal/pkg/config"
)

// RedisCache - implementation of Cache interface
type RedisCache struct {
	ctx    context.Context
	client *redis.Client
	config *config.Config
}

// InitRedisCache - create new instance of RedisCache
// host and port - connection to Redis instance
func InitRedisCache(ctx context.Context, conf *config.Config) (*RedisCache, error) {
	log.Println("Redis: ", conf.RedisAddress)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// check connection by setting test value
	err := rdb.Set(ctx, "key", "value", 0).Err()

	return &RedisCache{
		ctx:    ctx,
		client: rdb,
		config: conf,
	}, err
}

// Add - add rand value with expiration (in seconds) to cache
func (c *RedisCache) Add(key string) error {
	return c.client.Set(c.ctx, key, "value", c.config.RedisExpDur()).Err()
}

// Get - check existence of int key in cache
func (c *RedisCache) Get(key string) (bool, error) {
	val, err := c.client.Get(c.ctx, key).Result()
	return val != "", err
}

// Delete - delete key from cache
func (c *RedisCache) Delete(key string) {
	c.client.Del(c.ctx, key)
}
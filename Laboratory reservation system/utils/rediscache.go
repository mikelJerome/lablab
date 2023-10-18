package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {

	// 使用指定的选项创建一个新的 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果有的话
		DB:       0,                // Redis 数据库编号
		// 连接池设置
		PoolSize: 10, // 连接池中保持的连接数
	})
	red := RedisClient{
		client: client,
	}
	return &red, nil
}

func (redis *RedisClient) SeCachet(key string, value interface{}) {
	redis.client.Set(redis.client.Context(), key, value, 0)
}

// SGet 通过键从 Redis 中获取字符串值
func (redis *RedisClient) GetCache(key string) (string, error) {
	code, _ := redis.client.Get(context.Background(), key).Result()
	return code, nil
}

package utils

import (
	"github.com/go-redis/redis"
	"time"
)

var Redis *RedisClient

type RedisClient struct {
	*redis.Client
}

func NewRedisClient() error {
	if Redis != nil {
		return nil
	}
	// 使用指定的选项创建一个新的 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果有的话
		DB:       0,                // Redis 数据库编号
		// 连接池设置
		PoolSize: 10, // 连接池中保持的连接数
	})
	Redis = &RedisClient{
		client,
	}
	return nil
}

// set 存储键值
func (redis *RedisClient) SeCachet(key string, value interface{}, expiretiem time.Duration) error {
	//当我们在Go代码中执行一个Redis命令时（例如使用redis.Client的Set方法），返回的结果是一个redis.StatusCmd类型。通过redis.StatusCmd，我们可以获取执行命令的结果、错误信息以及其他相关的元数据。
	//
	//以下是一些常用的redis.StatusCmd方法和属性：
	//
	//Val()：获取命令执行的结果值。
	//Err()：获取执行命令期间发生的错误。
	//String() string：获取命令执行结果的字符串表示形式。
	//Result() (string, error)：获取命令执行的结果和错误，以元组形式返回。
	return redis.Set(key, value, expiretiem).Err()
}

// SGet 通过键从 Redis 中获取字符串值
func (redis *RedisClient) GetCache(key string) string {
	return redis.Get(key).String()
}

// 删除键值

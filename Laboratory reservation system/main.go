package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Redis *RedisClient

// RedisClient 扩展了 redis.Client 并添加了额外的功能
type RedisClient struct {
	*redis.Client
}

// NewRedisClient 初始化 Redis 客户端
func NewRedisClient() error {
	if Redis != nil {
		return nil
	}

	// 使用指定的选项创建一个新的 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果有的话
		DB:       1,                // Redis 数据库编号

		// 连接池设置
		PoolSize: 10, // 连接池中保持的连接数

		// 超时设置
		DialTimeout:  5 * time.Second, // 连接建立超时时间
		ReadTimeout:  3 * time.Second, // 读取超时时间，-1 表示无超时
		WriteTimeout: 3 * time.Second, // 写入超时时间，-1 表示无超时
		PoolTimeout:  4 * time.Second, // 从连接池获取连接的最大等待时间

		// 空闲连接检查设置
		IdleCheckFrequency: 60 * time.Second, // 空闲连接检查的频率
		IdleTimeout:        5 * time.Minute,  // 空闲连接的超时时间，-1 表示无超时
		MaxConnAge:         0 * time.Second,  // 连接的最大存活时间，0 表示无最大存活时间

		// 命令重试设置
		MaxRetries:      0,                      // 命令执行失败时的最大重试次数，0 表示不重试
		MinRetryBackoff: 8 * time.Millisecond,   // 重试间隔时间的最小值，-1 表示无间隔
		MaxRetryBackoff: 512 * time.Millisecond, // 重试间隔时间的最大值，-1 表示无间隔

		// 在需要从连接池获取连接时，如果连接池需要创建新连接，则调用此钩子函数
		OnConnect: func(conn *redis.Conn) error {
			fmt.Printf("创建新的连接：%v\n", conn)
			return nil
		},
	})

	// 向 Redis 服务器发送 Ping 命令以测试连接
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	// 将全局的 Redis 变量设置为新创建的 Redis 客户端
	Redis = &RedisClient{client}
	return nil
}

// 在导入包时初始化 Redis 客户端
func init() {
	err := NewRedisClient()
	if err != nil {
		fmt.Println("连接 Redis 客户端失败")
	}
}

// SSet 在 Redis 中设置一个字符串值，可选设置过期时间（默认为 24 小时）
func (redis *RedisClient) SSet(key string, value interface{}) *redis.StatusCmd {
	return redis.Set(key, value, 24*time.Hour)
}

// SGet 通过键从 Redis 中获取字符串值
func (redis *RedisClient) SGet(key string) string {
	return redis.Get(key).String()
}

// Close 关闭 Redis 客户端
func (redis *RedisClient) Close() {
	redis.Close()
}

// GetRedisClient 返回 Redis 客户端，如果未初始化，则进行初始化
func GetRedisClient() (*RedisClient, error) {
	if Redis == nil {
		err := NewRedisClient()
		if err != nil {
			return nil, err
		}
		return Redis, nil
	}
	return Redis, nil
}

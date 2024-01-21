package utils

import (
	"context"
	"github.com/go-redis/redis/v8" // 更新为 v8 版本
	"go.uber.org/zap"
	"lab_sys/global"
	"time"
	//redislock "github.com/jefferyjob/go-redislock"
)

// RedisClient 扩展了 redis.Client 并添加了额外的功能
//type RedisClient struct {
//	*redis.Client
//}

// NewRedisClient 初始化 Redis 客户端
func NewRedisClient() error {
	if global.Redis != nil {
		return nil
	}

	ctx := context.Background() // 创建一个 context

	// 使用指定的选项创建一个新的 Redis 客户端
	client := redis.NewClient(&redis.Options{
		// ... 其他选项保持不变 ...
		Addr:     "127.0.0.1:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果有的话
		DB:       0,                // Redis 数据库编号
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
		MaxRetries:      3,                      // 命令执行失败时的最大重试次数，0 表示不重试
		MinRetryBackoff: 8 * time.Millisecond,   // 重试间隔时间的最小值，-1 表示无间隔
		MaxRetryBackoff: 512 * time.Millisecond, // 重试间隔时间的最大值，-1 表示无间隔

	})

	// 向 Redis 服务器发送 Ping 命令以测试连接
	_, err := client.Ping(ctx).Result() // 使用 context
	if err != nil {
		zap.L().Error("测试不成功", zap.Error(err))
		return err
	}

	// 将全局的 Redis 变量设置为新创建的 Redis 客户端

	global.Redis = client

	return err
}

// 在导入包时初始化 Redis 客户端
//func init() {
//	err := NewRedisClient()
//	if err != nil {
//		fmt.Println("连接 Redis 客户端失败")
//	}
//}

// SSet 在 Redis 中设置一个字符串值，可选设置过期时间（默认为 24 小时）
//func (r RedisClient) SSet(ctx context.Context, key string, value interface{}) *redis.StatusCmd {
//	//ctx := context.Background()                 // 创建 context
//	return r.Set(ctx, key, value, 24*time.Hour) // 使用 context
//}

// SGet 通过键从 Redis 中获取字符串值
//func (r *RedisClient) SGet(ctx context.Context, key string) (string, error) {
//	//ctx := context.Background()     // 创建 context
//	return r.Get(ctx, key).Result() // 使用 context
//}

// Close 关闭 Redis 客户端
//func (r *RedisClient) Close() {
//	r.Client.Close()
//	//ctx := context.Background() // 创建 context
//}

// GetRedisClient 返回 Redis 客户端，如果未初始化，则进行初始化
func GetRedisClient() (*redis.Client, error) {
	if global.Redis == nil {
		err := NewRedisClient()
		if err != nil {
			return nil, err
		}
	}
	return global.Redis, nil
}

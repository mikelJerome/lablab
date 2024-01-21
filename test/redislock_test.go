package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	redislock "github.com/jefferyjob/go-redislock"
	"testing"
)

func TestRedislock(t *testing.T) {
	// 创建Redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 创建上下文
	ctx := context.Background()

	// 创建分布式锁
	lock := redislock.New(ctx, redisClient, "liuan")

	// 加锁
	err := lock.Lock()
	if err != nil {
		fmt.Println("Failed to acquire lock:", err)
		return
	}

	// 执行业务逻辑
	fmt.Println("Lock acquired. Performing critical section.")

	// 解锁
	err = lock.UnLock()
	if err != nil {
		fmt.Println("Failed to release lock:", err)
		return
	}

	fmt.Println("Lock released.")
}

package initialize

import (
	"context"
	"fmt"
	"lab_sys/global"
	"strconv"
	"testing"
	"time"
)

func TestInitRedis(t *testing.T) {
	InitRedis()
	ctx := context.Background()
	global.Redis.Set(ctx, strconv.Itoa(1111)+"111", "fdffdfdsfsfa", 10*time.Hour)
	result := global.Redis.Get(ctx, strconv.Itoa(1111)+"111")
	verifycode := result.Val()
	fmt.Printf("%q\n", verifycode)
}

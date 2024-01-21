package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"lab_sys/config"
	//redislock "github.com/jefferyjob/go-redislock"
)

var (
	DB       *gorm.DB
	Settings config.ServerConfig
	Lg       *zap.Logger
	Redis    *redis.Client
	//Redislock redislock.RedisLockInter
)

const (
	DefaultSender    = "262396235@qq.com"
	DefaultSPassword = "lbmdohlulybfcacf"
	DefaultSMTPAddr  = "smtp.qq.com"
	DefaultSMTPPort  = 25
	DefaultMailTitle = "您好,请查收实验室验证码。"
)

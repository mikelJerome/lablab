package model

import (
	"gorm.io/gorm"
	"time"
)

/*
用户 ID（userID）
用户名（username）
密码（password）
手机号码*（Phone）
器材id
*/

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Identity uint //防止并发出现问题，uuid
	Username string
	Password string
	Phone    string `valid:"matches(^1[1-9]{1}\\d{9}$)"`
	Email    string
}

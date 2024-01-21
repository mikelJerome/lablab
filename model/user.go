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
	Username string `gorm:"type:varchar(255)" form:"username" binding:"required"`
	Password string `valid:"matches{8,}$" form:"password" binding:"required"` // 废弃字段
	Phone    string `valid:"matches(^1[1-9]{1}\\d{9}$)" form:"phone" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
}

func (User) TableName() string {
	return "user"
}

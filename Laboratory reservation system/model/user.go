package model

/*
用户 ID（userID）
用户名（username）
密码（password）
手机号码*（Phone）
器材id
*/

//https://juejin.cn/post/696584936159916852088

type User struct {
	UserID   uint `gorm:"primaryKey"`
	Username string
	Password string
	Phone    string `valid:"matches(^1[1-9]{1}\\d{9}$)"`
	Email    string
}

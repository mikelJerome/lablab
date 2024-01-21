package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lab_sys/database"
	"lab_sys/global"
	"lab_sys/model"
	"net/http"
)

// Register 处理用户注册的请求
func Register(c *gin.Context) {
	var user model.User
	// 绑定请求参数到 user 结构体
	if err := c.ShouldBind(&user); err != nil {
		zap.L().Error("请求参数绑定失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数不正确: "})
		return
	}

	// 验证码校验逻辑
	verifycode := c.PostForm("verifycode")
	//fmt.Print(verifycode)
	if err := checkVerifyCode(user.Email, verifycode, c); err != nil {
		zap.L().Error("校验验证码失败")
		fmt.Println(c.Request.PostForm)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否已存在,通过  用户  手机号码查询
	if _, err := database.FindPhone(user.Phone); err == nil {
		zap.L().Error("尝试注册已经存在的用户", zap.Error(err))
		c.JSON(http.StatusConflict, gin.H{"message": "用户已经存在"})
		return
	}

	// 创建用户 ,调用数据库方法创建用户。
	user1, err := database.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败: " + err.Error()})
		return
	}

	// 注册成功，返回用户信息
	c.JSON(http.StatusOK, gin.H{"data": user1})
}

// checkVerifyCode 校验验证码是否正确
func checkVerifyCode(email, code string, c context.Context) error {
	// 实现保持不变...
	result := global.Redis.Get(c, "Email"+" "+email)
	verifycoderedis, _ := result.Result()
	//strings.TrimSpace(verifycoderedis)
	//fmt.Print(verifycoderedis)
	//if err != nil {
	//	return err
	//}
	if verifycoderedis == "" || code != verifycoderedis {
		fmt.Printf("%q\n", verifycoderedis)
		fmt.Printf("%q\n", code)
		return fmt.Errorf("验证码错误或者已过期")
	}
	return nil
}

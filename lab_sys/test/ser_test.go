package test

import (
	"lab_sys/lab_sys/utils"
	"testing"
	"time"
)

func TestVerification(t *testing.T) {
	qquser := "example@qq.com"
	code := "123456"

	// 设置验证码
	utils.SetVerification(qquser, code)

	// 检查验证码是否有效
	valid := utils.Verify(qquser, code)
	if !valid {
		t.Errorf("Expected verification code to be valid, but it wasn't")
	}

	// 检查过期验证码是否无效
	expiredTime := time.Now().Add(-1 * time.Minute)
	expiredCode := utils.VerificationCode{
		Code:       code,
		ExpireTime: expiredTime,
	}
	utils.SetVerification(qquser, expiredCode.Code)
	valid = utils.Verify(qquser, expiredCode.Code)
	if valid {
		t.Errorf("Expected expired verification code to be invalid, but it was valid")
	}
}

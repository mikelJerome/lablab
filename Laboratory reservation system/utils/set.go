package utils

import (
	"time"
)

// var c =sync.Mutex{}
type VerificationCode struct {
	Code       string
	ExpireTime time.Time
}

var verificationcodes map[string][]VerificationCode

func SetVerification(qquser string, code string) {
	expireTime := time.Now().Add(1 * time.Minute)
	verificationcode := VerificationCode{
		code,
		expireTime,
	}
	// 添加验证
	verificationcodes[qquser] = append(verificationcodes[qquser], verificationcode)
	// 统计过期验证码
	Verify(qquser, verificationcode)

}
func Verify(qquser string, verifcationcode VerificationCode) {
	currentTime := time.Now()
	for i := 0; i < len(verificationcodes); i++ {
		if verifcationcode.ExpireTime.Before(currentTime) {
			verificationcodes[qquser] = append(verificationcodes[qquser][:i], verificationcodes[qquser][i+1:]...)
		} else {
			i++
		}
	}
}

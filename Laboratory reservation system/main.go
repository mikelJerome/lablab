package main

import (
	"fmt"
	"time"
)

type VerificationCode struct {
	Code       string
	ExpiryTime time.Time
}

func main() {
	verificationCodes := make(map[string][]VerificationCode)

	// 生成一个验证码
	expiryTime := time.Now().Add(time.Minute * 5)
	verificationCode := VerificationCode{
		Code:       "your_verification_code",
		ExpiryTime: expiryTime,
	}

	// 将验证码添加到切片中
	verificationCodes["user1"] = append(verificationCodes["user1"], verificationCode)

	// 验证验证码
	enteredCode := "your_entered_code"
	currentTime := time.Now()

	codes := verificationCodes["user1"]
	for i := 0; i < len(codes); i++ {
		if codes[i].Code == enteredCode && !codes[i].ExpiryTime.Before(currentTime) {
			// 验证码匹配且未过期
			fmt.Println("验证码正确，并且未过期")
			break
		}
	}
}

package utils

/*
本来打算不用redis单纯的使用goalng做一个简单的缓存
但是这个代码好像还挺复杂的，复杂的是后面的逻辑处理复杂
谁看到这个项目对这个感兴趣可以帮我完善一下，我要开始使用redis了
代码好不容易写出来，舍不得删除了，就这样吧
这个文件不会做任何调用
*/

//import (
//	"sync"
//	"time"
//)
//
//type VerificationCode struct {
//	Code       string
//	ExpireTime time.Time
//}
//
//var verificationcodes map[string]VerificationCode
//var mutex sync.Mutex
//
//func SetVerification(qquser string, code string) {
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	expireTime := time.Now().Add(1 * time.Minute)
//	verificationcode := VerificationCode{
//		Code:       code,
//		ExpireTime: expireTime,
//	}
//
//	// 添加或更新验证码
//	verificationcodes[qquser] = verificationcode
//}
//
//func Verify(qquser string, verificationcode string) bool {
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	// 检查验证码是否存在
//	if code, ok := verificationcodes[qquser]; ok {
//		// 检查验证码是否过期
//		if code.ExpireTime.After(time.Now()) {
//			// 进行验证
//			return code.Code == verificationcode
//		}
//	}
//
//	return false
//}

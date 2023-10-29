package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

func Email(email string) string {
	var mailConf MailboxConf
	mailConf.Title = "验证"
	//这里支持群发，只需填写多个人的邮箱即可，我这里发送人使用的是QQ邮箱，所以接收人也必须都要是
	//QQ邮箱
	mailConf.RecipientList = email
	mailConf.Sender = `262396235@qq.com`

	mailConf.SPassword = "lbmdohlulybfcacf"
	mailConf.SMTPAddr = "smtp.qq.com"
	//下面是官方邮箱提供的SMTP服务地址和端口
	// QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	// 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	// 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	// 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	// 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）

	mailConf.SMTPPort = 25

	//产生六位数验证码
	verify := RandNumber(5)
	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, verify)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, mailConf.RecipientList)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	_ = gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	//if err != nil {
	//	b, _ := json.Marshal("验证码返回失败，请输入正确的邮箱")
	//	fmt.Printf(string(b))
	//	return
	//}
	//b, _ := json.Marshal("验证码返回成功，输入的邮箱有效")
	//fmt.Printf(string(b))
	return verify
}

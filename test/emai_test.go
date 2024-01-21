package test

import (
	"fmt"
	"lab_sys/global"
	"lab_sys/utils"
	"testing"
)

func Test_Email(t *testing.T) {

	mailConf := utils.MailboxConf{
		Title:         global.DefaultMailTitle,
		RecipientList: "2353208816@qq.com",
		Sender:        global.DefaultSender,

		SPassword: global.DefaultSPassword,
		SMTPAddr:  global.DefaultSMTPAddr,
		SMTPPort:  25,
	}
	fmt.Printf("%s", mailConf.Email())

}

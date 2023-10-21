package main

import (
	"errors"
	"lab_sys/lab_sys/global"
	"lab_sys/lab_sys/model"
)

// 获取用户列表
func GetUserList() ([]*model.User, error) {
	var list []*model.User
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("get user list fail")
	}
	return list, nil
}

//用户查询
func FindUser(name string) (*model.User, error) {
	user:=model.user
}

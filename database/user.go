package database

import (
	"errors"
	"go.uber.org/zap"
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

// 电话查询（注册用）
// 注册时保证电话的唯一
func FindPhone(phone string) (*model.User, error) {
	user := model.User{}
	if tx := global.DB.Where("phone = ?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("phone is exist")
	}
	return &user, nil
}

// 电话查询
func FindUserByPhone(phone string) (*model.User, error) {
	user := model.User{}
	if tx := global.DB.Where("phone = ?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("the user is not queried")
	}
	return &user, nil
}

// 用户名查询
func FindUserByName(name string) (*model.User, error) {
	user := model.User{}
	if tx := global.DB.Where("username = ?", name).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("the user is not queried")
	}
	return &user, nil
}

// 创建用户
func CreateUser(user model.User) (*model.User, error) {
	if tx := global.DB.Create(&user); tx.RowsAffected == 0 {
		zap.S().Info("create failed")
		return nil, errors.New("create failed")
	}
	return &user, nil
}

// 更新用户
func UpdateUser(user model.User) (*model.User, error) {
	tx := global.DB.Model(&user).Updates(model.User{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("update failed")
		return nil, errors.New("update failed")
	}
	return &user, nil
}

// 删除用户
func DeleteUser(user model.User) error {
	tx := global.DB.Delete(&user)
	if tx.RowsAffected == 0 {
		zap.S().Info("delete failed")
		return errors.New("delete failed")
	}
	return nil
}

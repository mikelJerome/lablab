package test

import (
	"lab_sys/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSql(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/lab?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Equipment{})
	db.AutoMigrate(&model.Reservation{})

}

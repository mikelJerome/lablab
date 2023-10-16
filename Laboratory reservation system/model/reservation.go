package model

import "time"

/*
预约 ID（reservationID）
用户 ID（userID）
器材 ID（equipmentID）
预约时间（reservationTime）
预约状态（status，如待使用、已使用、已取消
*/

// 预约记录表

type Reservation struct {
	ReservationID   uint `gorm:"primaryKey"`
	Username        string
	EquipmentID     uint
	ReservationTime time.Time `gorm:"type:datetime"`
	Status          string
}

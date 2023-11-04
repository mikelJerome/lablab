package database

import (
	"errors"
	"lab_sys/global"
	"lab_sys/model"

	"go.uber.org/zap"
)

// 获取预约列表
func GetReservationList() ([]*model.Reservation, error) {
	var list []*model.Reservation
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("get reservation list failed")
	}
	return list, nil
}

// 根据预约id查找记录
func FindReservation(reservationId string) (*model.Reservation, error) {
	reservation := model.Reservation{}
	if tx := global.DB.Where("identity =? ", reservationId).First(&reservation); tx.RowsAffected == 0 {
		return nil, errors.New("the reservation is not queried")
	}
	return &reservation, nil
}

// 根据用户查找预约记录
func FindReservationByUser(name string) (*model.Reservation, error) {
	reservation := model.Reservation{}
	tx := global.DB.First(&reservation).Joins(" join user on reservation.user_id=user.identity").Where("user.username = ?", name)
	if tx.RowsAffected == 0 {
		return nil, errors.New("the reservation is not queried")
	}
	return &reservation, nil
}

// 根据器材查找预约记录
func FindReservationByEquip(equip string) (*model.Reservation, error) {
	reservation := model.Reservation{}
	tx := global.DB.First(&reservation).Joins(" joie equipment on reservation.equipment_id=equipment.identity").Where("equipment.name = ?", equip)
	if tx.RowsAffected == 0 {
		return nil, errors.New("the reservation is not queried")
	}
	return &reservation, nil
}

// 创建预约
func CreateReservation(reservation model.Reservation) (*model.Reservation, error) {
	tx := global.DB.Create(&reservation)
	if tx.RowsAffected == 0 {
		zap.S().Info("create reservation failed")
		return nil, errors.New("create reservation failed")
	}
	return &reservation, nil
}

// 更新预约
func UpdateReservation(reservation model.Reservation) (*model.Reservation, error) {
	tx := global.DB.Updates(model.Reservation{
		UserID:      reservation.UserID,
		EquipmentID: reservation.EquipmentID,
		StartTime:   reservation.StartTime,
		OverTime:    reservation.OverTime,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("update failed")
		return nil, errors.New("update failed")
	}
	return &reservation, nil
}

// 删除预约
func DeleteReservation(reservation model.Reservation) error {
	tx := global.DB.Delete(&reservation)
	if tx.RowsAffected == 0 {
		zap.S().Info("delete failed")
		return errors.New("delete failed")
	}
	return nil
}

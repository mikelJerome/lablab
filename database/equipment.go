package database

import (
	"errors"
	"lab_sys/global"

	"lab_sys/model"

	"go.uber.org/zap"
)

// 获取器材列表
func GetEquipmentList() ([]*model.Equipment, error) {
	var list []*model.Equipment
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("get equipment list failed")
	}
	return list, nil
}

// 创建器材
func CreateEquipment(equip model.Equipment) (*model.Equipment, error) {
	tx := global.DB.Create(&equip)
	if tx.RowsAffected == 0 {
		zap.S().Info("create failed")
		return nil, errors.New("create failed")
	}
	return &equip, nil
}

// 查询器材
func FindEquipmentByName(name string) (*model.Equipment, error) {
	equip := model.Equipment{}
	if tx := global.DB.Where("name = ?", name).First(&equip); tx.RowsAffected == 0 {
		return nil, errors.New("find failed")
	}
	return &equip, nil
}

// 更新器材
func UpdateEquipment(equip model.Equipment) (*model.Equipment, error) {
	tx := global.DB.Updates(model.Equipment{

		Name:  equip.Name,
		Type:  equip.Type,
		Stock: equip.Stock,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("update failed")
		return nil, errors.New("update failed")
	}
	return &equip, nil
}

// 删除器材
func DeleteEquip(equip model.Equipment) error {
	tx := global.DB.Delete(&equip)
	if tx.RowsAffected == 0 {
		zap.S().Info("delete failed")
		return errors.New("delete failed")
	}
	return nil
}

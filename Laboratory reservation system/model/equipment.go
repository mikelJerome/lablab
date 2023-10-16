package model

/*

器材 ID（equipmentID）
器材名称（name）
*/

type Equipment struct {
	EquipmentID uint `gorm:"primaryKey"`
	Name        string
	Type        string
	Stock       int
}

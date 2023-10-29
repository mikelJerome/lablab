package model

/*

器材 ID（equipmentID）
器材名称（name）
*/

type Equipment struct {
	Model
	Name  string `gorm:"type:varchar(255)"`
	Type  string
	Stock int
}

func (Equipment) TableName() string {
	return "equipment"
}

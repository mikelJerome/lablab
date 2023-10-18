package model

/*

器材 ID（equipmentID）
器材名称（name）
*/

type Equipment struct {
	Model
	Identity uint
	Name     string
	Type     string
	Stock    int
}

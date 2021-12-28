package models


type Role struct {
	Id       string     `json:"id" xorm:"id"`
	Name     string `json:"name" xorm:"name"`
	IsSuper  bool   `json:"is_super" xorm:"is_super"`
	Status   bool   `json:"status" xorm:"status"`
	Permt    []*Menu
}


func(a *Role) TableName() string {
	return "role"
}
package models


type Role struct {
	Name     string `json:"name" xorm:"name"`
	IsSuper  bool   `json:"is_super" xorm:"is_super"`
	Enable   bool   `json:"enable" xorm:"enable"`
	Permt   []*Menu
}

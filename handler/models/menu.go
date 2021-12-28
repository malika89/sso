package models


type Menu struct {
	Id           string     `json:"id" xorm:"id"`
	Name         string     `json:"name" xorm:"name"`
	Status       bool       `json:"status" xorm:"status"`
	Category     string     `json:"category" xorm:"category"`  //page,button
	Application  Application
	Parent       *Menu      `json:"id" xorm:"id"`
}

func(a *Menu) TableName() string {
	return "menu"
}
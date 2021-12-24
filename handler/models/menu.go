package models


type Menu struct {
	Id           string     `json:"id" xorm:"id"`
	Name         string     `json:"name" xorm:"name"`
	Enable       bool       `json:"enable" xorm:"enable"`
	Category     string     `json:"category" xorm:"category"`  //page,button
	Application  Application
	Parent       *Menu      `json:"id" xorm:"id"`
}

package models

type Application struct {
	Id       string  `json:"id" xorm:"id"`
	Name     string  `json:"name" xorm:"name"`
	Domain   string  `json:"domain" xorm:"domain"`
	Enable   bool    `json:"enable" xorm:"enable"`
	Chargeman User   `json:"chargeman" xorm:"chargeman"`
}



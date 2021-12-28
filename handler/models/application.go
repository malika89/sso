package models

type Application struct {
	Id       string  `json:"id" xorm:"id"`
	Name     string  `json:"name" xorm:"name"`
	Domain   string  `json:"domain" xorm:"domain"`
	Status   bool    `json:"status" xorm:"status"`
	Chargeman User   `json:"chargeman" xorm:"chargeman"`
}


func(a *Application) TableName() string {
	return "application"
}
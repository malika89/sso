package models

import "time"

type User struct {
	Id       string     `json:"id" xorm:"id"`
	Name       string    `json:"name" xorm:"name"`
	JoinTime   time.Time `json:"join_time" xorm:"join_time"`
	LoginIp    string    `json:"login_ip" xorm:"login_ip"`
	Status     string    `json:"status" xorm:"status"`
	Dept       Dept      `json:"dept" xorm:"dept"`
	Role       []*Role   `json:"role" xorm:"role"`
}

func(a *User) TableName() string {
	return "user"
}

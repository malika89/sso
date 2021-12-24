package models

import "time"

type User struct {
	Name       string    `json:"name" xorm:"name"`
	JoinTime   time.Time `json:"join_time" xorm:"join_time"`
	LoginIp    string    `json:"login_ip" xorm:"login_ip"`
	Enable     string    `json:"enable" xorm:"enable"`
	Dept       Dept      `json:"dept" xorm:"dept"`
	Role       []*Role   `json:"role" xorm:"role"`
}
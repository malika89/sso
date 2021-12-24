package models

type Organization struct {
	Id      string   `json:"id" xorm:"id"`
	Name    string   `json:"name" xorm:"name"`
	Address string   `json:"address" xorm:"address"`
	Enable  bool     `json:"enable" xorm:"enable"`
	Dept    []*Dept  `json:"dept" xorm:"dept"`
}

type Dept struct {
	Id          string   `json:"id" xorm:"id"`
	Name        string   `json:"name" xorm:"name"`
	Code        string   `json:"code" xorm:"code"`
	Enable      bool     `json:"enable" xorm:"enable"`
	EmployeeNo  string   `json:"employee_no" xorm:"employee_no"`
	SetUpTime   string   `json:"setup_time" xorm:"setup_time"`
	Level       int      `json:"level" xorm:"level"`               //等级
	Parent      *Dept    `json:"parent_id" xorm:"parent_id"`
}


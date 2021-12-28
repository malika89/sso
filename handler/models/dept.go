package models

type Organization struct {
	Id      string   `json:"id" xorm:"id"`
	Name    string   `json:"name" xorm:"name"`
	Address string   `json:"address" xorm:"address"`
	Status  bool     `json:"status" xorm:"status"`
	Dept    []*Dept  `json:"dept" xorm:"dept"`
}

func(a *Organization) TableName() string {
	return "organization"
}

type Dept struct {
	Id          string   `json:"id" xorm:"id"`
	Name        string   `json:"name" xorm:"name"`
	Code        string   `json:"code" xorm:"code"`
	Status      bool     `json:"status" xorm:"status"`
	EmployeeNo  string   `json:"employee_no" xorm:"employee_no"`
	SetUpTime   string   `json:"setup_time" xorm:"setup_time"`
	Slevel       int     `json:"slevel" xorm:"slevel"`               //等级
	Parent      *Dept    `json:"parent_id" xorm:"parent_id"`
}

func(a *Dept) TableName() string {
	return "dept"
}



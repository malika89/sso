package models

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	//"sso/conf"
	"github.com/goinggo/mapstructure"
)

var DBX *xorm.Engine

func init()  {
	initDB()

}
func initDB() error {
	var err error
	//engine, err = xorm.NewEngine(conf.Client.GetValue("db.driver"), conf.Client.GetValue("db.host"))
	DBX, err = xorm.NewEngine("mysql","root:PACloud@20!^@tcp(127.0.0.1:3306)/sso?charset=utf8")
	if err !=nil{
		return err
	}
	fmt.Println("init database success")
	return nil
}

type BaseModel struct {
	Table string
	Model interface{}  //struct结构体
}

func(t *BaseModel) TableName() string {
	return "base_model"
}

func(t *BaseModel) GetSession() *xorm.Session {
	if t.Table == "" {
		t.Table =t.TableName()
	}
	dbSession := DBX.Table(t.Table)
	return dbSession

}
func(t *BaseModel) Exists(keyword string,value string ) (bool,error) {

	infStruct := t.Model
	dbSession := t.GetSession()
	filterMap := map[string]string{keyword:value}
	if err := mapstructure.Decode(filterMap, infStruct); err != nil {
		return false,err
	}
	return dbSession.Exist(&infStruct)
}

func(t *BaseModel) Query(keyword string,value string ) ([]interface{},error) {
	var infLst []interface{}
	dbSession := t.GetSession()
	err := dbSession.Where(fmt.Sprintf("%s=%s",keyword,value)).Find(&infLst)
	return infLst,err
}


func(t *BaseModel) Update() error {
	var modelMap map[string]interface{}
	byteValue,_ :=json.Marshal(t.Model)
	json.Unmarshal(byteValue,modelMap)

	dbSession := t.GetSession()
	_ ,err := dbSession.ID(modelMap["Id"].(string)).Update(&t.Model)
	return err
}

func(t *BaseModel) Delete() error {
	var modelMap map[string]interface{}
	byteValue,_ :=json.Marshal(t.Model)
	json.Unmarshal(byteValue,modelMap)
	dbSession := t.GetSession()
	_ ,err := dbSession.ID(modelMap["Id"].(string)).Delete(&t.Model)
	return err
}


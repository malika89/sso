package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/goinggo/mapstructure"
	"sso/conf"
)

var DBX *xorm.Engine
var cfg =conf.Conf

func init() {
	initDB()
}

func initDB() error {
	if DBX ==nil{
		var err error
		//engine, err = xorm.NewEngine(conf.Client.GetValue("db.driver"), conf.Client.GetValue("db.host"))
		connStr := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
			"127.0.0.1", "5432","root", "postgres", "xops_sso")
		DBX, err = xorm.NewEngine("postgres",connStr)
		if err !=nil{
			return err
		}
		DBX.ShowSQL(true)
		fmt.Println("init database success")
	}
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
	if DBX == nil{
		initDB()
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

func(t *BaseModel) Query(keyword string,value string ) ([]map[string]string,error) {
	dbSession := t.GetSession()
	infLst,err := dbSession.Where(keyword +" = ?",value).QueryString()
	return infLst,err
}

func(t *BaseModel) Add() error {
	dbSession := t.GetSession()
	if _, err := dbSession.Insert(t.Model); err != nil {
		return err
	}
	return nil
}

func(t *BaseModel) Update() error {
	dbSession := t.GetSession()
	modelMap :=t.Model.(map[string]interface{})
	if _, err := dbSession.ID(modelMap["id"].(string)).Update(modelMap); err != nil {
		return err
	}
	return nil
}

//Delete 支持批量删除
func(t *BaseModel) Delete() error {
	dbSession := t.GetSession()
	if _, err := dbSession.Delete(t.Model); err != nil {
		return err
	}
	return nil
}

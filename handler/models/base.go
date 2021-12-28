package models

import (
	"encoding/json"
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/go-xorm/xorm"
	//"sso/conf"
	"github.com/goinggo/mapstructure"
)

var DBX *xorm.Engine

func init()  {
	initDB()

}
func initDB() error {
	if DBX ==nil{
		var err error
		//engine, err = xorm.NewEngine(conf.Client.GetValue("db.driver"), conf.Client.GetValue("db.host"))
		connStr := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
			"10.25.72.97", "5432","postgres", "postgres", "xops_sso")
		DBX, err = xorm.NewEngine("postgres",connStr)
		DBX.ShowSQL(true)
		if err !=nil{
			return err
		}
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


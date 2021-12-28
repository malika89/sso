package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"os"
	"sso/conf"
	"sso/handler"
	"sso/utils"
)

var(
	cfg = conf.Conf
)

func start() error {
	if utils.CheckLink(cfg.Server.Host, cfg.Server.Port) {
		return fmt.Errorf("address: %s:%d was binded", cfg.Server.Host, cfg.Server.Port)
	}
	router :=handler.InitRoute()

	addr := fmt.Sprintf(`%s:%d`,cfg.Server.Host,cfg.Server.Port)
	s := endless.NewServer(addr,router)
	err := s.ListenAndServe()
	if err !=nil{
		fmt.Printf("start system error:%v",err)
		return err
	}
	return nil
}


func main(){
	//Default将使用logger 和 recovery中间件
	if err :=start();err!=nil{
		os.Exit(2)
	}
}



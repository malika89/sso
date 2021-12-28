package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"sso/conf"
	"sso/handler/views"
	"sso/utils"
)

var(
	cfg = conf.Conf
)
func setRoute(r *gin.Engine) {
	r.Handle("GET","/user/search",views.QueryUserByNameHandler)
	r.Handle("POST","/user",views.UpdateUserHandler)
	r.Handle("PUT","/user",views.AddUserHandler)
	r.Handle("DELETE","/user",views.DeleteUserHandler)
}

func start() error {
	if utils.CheckLink(cfg.Server.Host, cfg.Server.Port) {
		return fmt.Errorf("address: %s:%d was binded", cfg.Server.Host, cfg.Server.Port)
	}
	r :=gin.New()
	r.Use(gin.Recovery())
	setRoute(r)

	addr := fmt.Sprintf(`%s:%d`,cfg.Server.Host,cfg.Server.Port)
	s := endless.NewServer(addr,r)
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



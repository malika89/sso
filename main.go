package main

import (
	"github.com/gin-gonic/gin"
	"sso/handler/views"
	//"sso/conf"
)

func main(){
	r :=gin.Default()
	r.Handle("GET","/user/search",views.QueryUserByNameHandler)
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{"message":"pong"})
	})
	r.Run("0.0.0.0:8087")
}



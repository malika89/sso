package handler

import (
	"github.com/gin-gonic/gin"
	"sso/handler/views"
)

func InitRoute() *gin.Engine{
	r :=gin.New()
	r.Use(gin.Recovery())
	r.Handle("GET","/user/search",views.QueryUserByNameHandler)
	r.Handle("POST","/user",views.AddUserHandler)
	r.Handle("PUT","/user",views.UpdateUserHandler)
	r.Handle("DELETE","/user",views.DeleteUserHandler)
	return r
}

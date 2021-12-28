package handler

import (
	"github.com/gin-gonic/gin"
	"sso/handler/views"
)

func InitRoute() *gin.Engine{
	r :=gin.New()
	r.Use(gin.Recovery())

	v1 :=r.Group("/api/v1")
	{
		v1.GET("/user/search",views.QueryUserByNameHandler)
	    v1.POST("/user",views.AddUserHandler)
	    v1.PUT("/user",views.UpdateUserHandler)
	    v1.DELETE("/user",views.DeleteUserHandler)
	}
	return r
}

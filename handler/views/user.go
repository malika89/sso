package views

import (
	"github.com/gin-gonic/gin"
	"sso/handler/services"
	"sso/utils"
)

func QueryUserByNameHandler(c *gin.Context) {
	// 获取参数
	name := c.Param("name")
	userService := services.GetService()
	userDao, err := userService.Query("name",name)
	if userDao == nil || err !=nil {
		utils.BadResponse(c,utils.Error, "请求数据库错误")
		return
	}
	utils.NormalResponse(c, utils.OK, "",utils.ResponseObject{Data: userDao})
}

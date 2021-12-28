package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sso/handler/services"
	"sso/utils"
)

func QueryUserByNameHandler(c *gin.Context) {
	// 获取参数
	name := c.Query("name")
	userService := services.GetService()
	userDao, err := userService.Query("name",name)
	if err !=nil {
		utils.BadResponse(c,utils.Error, fmt.Sprintf("请求数据库错误:%v",err))
		return
	}
	utils.NormalResponse(c, utils.OK, "success",userDao,len(userDao))
}

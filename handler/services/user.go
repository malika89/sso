package services

import (
	"sso/handler/models"
)

type service struct {
	models.BaseModel
}

func GetService() *service  {
	var serviceApp service
	serviceApp.Table = "users"
	serviceApp.Model = models.User{}
	return &serviceApp
}

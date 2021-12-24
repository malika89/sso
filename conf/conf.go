package conf

/* apollo接入*/

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/spf13/viper"

)

var Conf Config
var Client agollo.Client


type Config struct {
	Apollo Apollo
	Server Server
}

type Apollo struct {
	AppID string
	Cluster string
	IP string
	NamespaceName string
	IsBackupConfig bool
	Secret string
}

type Server struct {
	Host string
	Port int
	Env string
}


func loadLocalConf() error {
	viper.AddConfigPath("./conf.toml")
	viper.SetConfigType("toml")
	if err :=viper.ReadInConfig();err !=nil{
		return err
	}
	if err :=viper.Unmarshal(&Conf);err !=nil{
		return err
	}
	return nil
}

func initApollo() error{
	var err error

	loadLocalConf()
	fmt.Println("load local config from ..")
	appConfig := &config.AppConfig{
		AppID:          Conf.Apollo.AppID,
		Cluster:         Conf.Apollo.Cluster,
		IP:              Conf.Apollo.IP,
		NamespaceName:   Conf.Apollo.NamespaceName,
		IsBackupConfig:  Conf.Apollo.IsBackupConfig,
		Secret:          Conf.Apollo.Secret,
	}
	Client,err =agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return appConfig,nil

	})
	if err!=nil{
		return err
	}
	return nil
}


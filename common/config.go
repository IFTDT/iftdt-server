package common

import (
	"github.com/spf13/viper"
)

type Config struct {
	GIN_MODE   string
	ServerPort string
	JWTKey     string

	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string

	APPID string
}

var ENV Config

func init() {
	viper.SetConfigFile(".env")

	// 自动加载环境变量进行覆盖
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("you should run this commond `copy .env.example .env`")
	}

	viper.Unmarshal(&ENV)
}

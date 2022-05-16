package common

import (
	"fmt"

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

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// path for unit test
	viper.AddConfigPath("../")

	// viper.SetConfigFile(".env")

	// 自动加载环境变量进行覆盖
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("you should run this commond `copy .env.example .env`")
	}

	viper.Unmarshal(&ENV)
	fmt.Print(ENV)
}

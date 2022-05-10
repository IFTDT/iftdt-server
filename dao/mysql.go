package dao

import (
	"fmt"

	"github.com/iftdt/server/common"
	"github.com/iftdt/server/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func init() {
	// 示例：user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", common.ENV.DBUser, common.ENV.DBPassword, common.ENV.DBHost, common.ENV.DBPort, common.ENV.DBDatabase)
	logger.Logger.Print(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

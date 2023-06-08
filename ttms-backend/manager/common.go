/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/06/06 23:03
 */

package manager

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mognolia/internal/dao"
	"os"
	"time"
)

func init() {
	dsn := "root:zxz123456@tcp(127.0.0.1:3306)/gin-vue-ttms?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢阙值
			Colorful:      true,        //禁用彩色
			LogLevel:      logger.Info,
		})
	//全局模式
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}
	dao.Group.DB = DB
}

package core

import (
	"fmt"
	"go-blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {

	cf := global.Config
	log := global.Log
	//log := global.Log
	if cf.Mysql.Host == "" {
		//fmt.Println("未配置mysql, 取消gorm连接")
		log.Error("未配置mysql, 取消gorm连接")
	}

	dsn := cf.Mysql.DSN()

	var sqlLogger logger.Interface
	if cf.System.Env == "dev" {
		// 开发环境显示所有sql
		sqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		sqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: sqlLogger,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("[%s] mysql连接失败", dsn))
		log.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Duration(cf.Mysql.MaxIdealTime) * time.Hour) // 最大复用时间
	sqlDB.SetMaxOpenConns(cf.Mysql.MaxOpenConns)                               // 最大连接
	sqlDB.SetMaxIdleConns(cf.Mysql.MaxIdleConns)                               // 最大空闲连接

	//fmt.Println(dsn, "数据库连接成功")
	log.Info(dsn, "数据库连接成功")
	return db

}

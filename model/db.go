package model

import (
	"fmt"
	"gin-casbin/utils/load"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(setting *load.Setting) *gorm.DB {
	var DB *gorm.DB
	var err error
	if setting == nil {
		fmt.Println("数据库连接失败，无效的Setting")
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		setting.DbHost,
		setting.DbUser,
		setting.DbPassWord,
		setting.DbName,
		setting.DbPort,
		setting.DbTimeZone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		//Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接失败,请检查参数", err)
		return nil
	}
	// 数据库自动迁移
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("数据库创建失败", err)
		return nil
	}
	return DB
}

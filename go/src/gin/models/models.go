package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dsn string = "root:''@tcp(39.106.6.99:3306)/shop?charset=utf8mb4" // 数据库配置 root账号  0000密码  shop数据库  utf8mb4编码
)

//实现orm
func DB() *gorm.DB {
	// 连接mysql数据库，  gorm.Config有很多配置 例如前缀  后准 等等
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tp_",
			SingularTable: true,
		},
	})

	// 数据库连接出错 退出
	if err != nil {
		fmt.Println("数据库连接失败", err)
		panic(1)
	}

	// 返回数据库实例
	return db
}

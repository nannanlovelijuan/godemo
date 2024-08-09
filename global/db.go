package global

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type DB struct {
// 	app *Application
// }

func InitDB(app *Application) *gorm.DB {

	_, dsn := app.Config.GetKv("db.dsn")
	fmt.Println(dsn)

	//根据配置获取数据库地址
	// dsn := "root:123456@tcp(192.168.12.118:3306)/ezp-bigdata?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// sqlite 启动报错，需要c+的编译环境
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	return db
}

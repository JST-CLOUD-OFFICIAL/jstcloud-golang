package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// 初始化数据库
func Init() error {
	source := "%s:%s@tcp(%s:%d)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&parseTime=true"
	user := os.Getenv("db_user")
	pwd := os.Getenv("db_pass")
	addr := os.Getenv("db_addr")
	database := os.Getenv("db_name")
	port := 3306

	source = fmt.Sprintf(source, user, pwd, addr, port, database)
	fmt.Println("start init mysql with ", source)

	// 链接数据库
	db, err := gorm.Open(mysql.Open(source))

	if err != nil {
		fmt.Println("db open error, err=", err.Error())
		return err
	}

	dbIntanceTemp, err := db.DB()
	if err != nil {
		fmt.Println("db init error,err=", err.Error())
		return err
	}

	// 设置打开数据库连接的最大数量
	dbIntanceTemp.SetMaxOpenConns(200)

	// 设置了连接可复用的最大时间
	dbIntanceTemp.SetConnMaxLifetime(time.Hour)

	// 设置连接池中空闲连接的最大数量
	dbIntanceTemp.SetMaxIdleConns(50)

	dbInstance = db

	fmt.Println("db init finished", source)
	return nil
}

// Get ...
func Get() *gorm.DB {
	return dbInstance
}

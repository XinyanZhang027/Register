package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zhuce/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	/*
		driverName := "mysql"
		host := "localhost"
		port := "3306"
		database := "go_demo"
		username := "go_admin"
		password := "123456"
		charset := "utf8"
		args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			username,
			password,
			host,
			port,
			database,
			charset)
		db, err := gorm.Open(driverName, args)
	*/
	dsn := "go_admin:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

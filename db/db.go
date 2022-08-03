package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=myuser dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if db == nil {
		fmt.Println("DB nil")
		return
	}
	DB = db
}

func GetDB() *gorm.DB {
	if DB == nil {
		fmt.Println("DB nil")
	}
	return DB
}

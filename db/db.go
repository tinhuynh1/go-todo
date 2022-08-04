package db

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	host := viper.GetString("postgresql.host")
	user := viper.GetString("postgresql.user")
	password := viper.GetString("postgresql.password")
	dbname := viper.GetString("postgresql.database")
	port := viper.GetString("postgresql.port")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, port)
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

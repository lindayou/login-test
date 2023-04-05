package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"login-server/model"
)

func main() {
	dsn := "root:123456@tcp(192.168.174.128:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&model.User{})
}

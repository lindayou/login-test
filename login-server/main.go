package main

import (
	"fmt"
	"login-server/model"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	groot := os.Getenv("GOROOT")
	fmt.Println(groot)
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("this is dsn", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 生成表
	//db.AutoMigrate(&model.User{})

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var result model.User
		if err := db.Where("username = ?", user.Username).First(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		if user.Password != result.Password {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login success"})
	})

	if err := r.Run(":8090"); err != nil {
		fmt.Println(err)
	}
}

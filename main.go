package main

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-borad/board-backend/db"
	"github.com/simple-borad/board-backend/router"
	"github.com/simple-borad/board-backend/router/auth"
	"gorm.io/gorm"
)

func SetupRouter(connect *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/hello", router.GetHello)

	r.POST("/auth/login", auth.Login)

	return r
}

func main() {
	connect := db.Initalize()
	router := SetupRouter(connect)
	router.Run()
}

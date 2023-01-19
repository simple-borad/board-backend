package main

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-borad/board-backend/router"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", router.GetHello)
	return r
}

func main() {
	router := SetupRouter()
	router.Run()
}

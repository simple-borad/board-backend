package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

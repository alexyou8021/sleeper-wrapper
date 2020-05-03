package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/alexyou8021/sleeper-wrapper.git/controllers"
)

func Handler() gin.HandlerFunc {
	result := controllers.Controller()
	return result
}

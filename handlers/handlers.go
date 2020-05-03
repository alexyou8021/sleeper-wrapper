package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/alexyou8021/sleeper-wrapper.git/controllers"
)

func Handler(ctx *gin.Context) {
	username := ctx.Param("name")
	result := controllers.Controller(username)
	ctx.JSON(http.StatusOK, result)
}

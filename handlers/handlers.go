package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/alexyou8021/sleeper-wrapper.git/controllers"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"
)

func Handler(ctx *gin.Context) {
	username := ctx.Param("name")
	result := controllers.Controller(username)
	ctx.JSON(http.StatusOK, result)
}

func GetESPNTransactions(ctx *gin.Context) {
	id := ctx.Param("id")
	result := controllers.GetESPNTransactions(id)
	ctx.JSON(http.StatusOK, result)
}

func CreateSleeperPlayersTable(ctx *gin.Context) {
	db.RemakeSleeperPlayersTable()
	ctx.JSON(http.StatusOK, "success")
}

func CreateESPNPlayersTable(ctx *gin.Context) {
	db.RemakeESPNPlayersTable()
	ctx.JSON(http.StatusOK, "success")
}

func CreateStatsTable(ctx *gin.Context) {
	db.RemakeStatsTable()
	ctx.JSON(http.StatusOK, "success")
}

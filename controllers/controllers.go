package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Controller() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	}
}

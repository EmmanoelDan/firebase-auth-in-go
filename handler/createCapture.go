package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCaptureHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST /capture",
	})
}

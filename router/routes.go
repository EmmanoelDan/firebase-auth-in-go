package router

import (
	"EmmanoelDan/firebase-auth-in-go.git/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		v1.GET("/capture", handler.ShowCaptureHandler)

		v1.POST("/capture", handler.CreateCaptureHandler)
		v1.PUT("/capture", handler.UpdateCaptureHandler)
		v1.DELETE("/capture", handler.DeleteCaptureHandler)
		v1.GET("/captures", handler.ListCaptureHandler)
	}
}

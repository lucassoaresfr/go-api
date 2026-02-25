package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucassoaresfr/go-api.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitHandle()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/createOpening", handler.CreateOpeningHandler)

		v1.DELETE("/deleteOpening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}
}

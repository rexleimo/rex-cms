package routes

import (
	"github.com/gin-gonic/gin"
	"rexai.com/controllers"
	"rexai.com/middlewares"
)

type ImagePage struct {
}

func (page *ImagePage) New(r *gin.RouterGroup) {
	imageC := controllers.ImageController{}
	group := r.Group("/image")
	{
		group.GET("", imageC.List)
		group.GET("/:id", imageC.Show)
		group.POST("", middlewares.AdminMiddleWare(), imageC.Create)
		group.POST("/publish", middlewares.AdminMiddleWare(), imageC.Upload)
		group.PUT("/:id", middlewares.AdminMiddleWare(), imageC.Update)
		group.DELETE("/:id", middlewares.AdminMiddleWare(), imageC.Delete)
	}
}

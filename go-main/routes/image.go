package routes

import (
	"github.com/gin-gonic/gin"
	"rexai.com/controllers"
)

type ImagePage struct {
}

func (page *ImagePage) New(r *gin.Engine) {
	imageC := controllers.ImageController{}
	group := r.Group("/image")
	{
		group.GET("", imageC.List)
	}
}

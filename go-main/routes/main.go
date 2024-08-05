package routes

import "github.com/gin-gonic/gin"

type Router interface {
	New(r *gin.Engine)
}

func RegisterRoutes(r *gin.Engine) {
	imgController := &ImagePage{}
	imgController.New(r)
}

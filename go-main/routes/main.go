package routes

import "github.com/gin-gonic/gin"

type Router interface {
	New(r *gin.Engine)
}

func RegisterRoutes(r *gin.Engine) {
	imgController := &ImagePage{}
	adminController := &AdminPage{}
	api := r.Group("/api")
	{
		imgController.New(api)
		adminController.New(api)
	}

}

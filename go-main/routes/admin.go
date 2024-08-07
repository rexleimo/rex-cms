package routes

import (
	"github.com/gin-gonic/gin"
	"rexai.com/controllers"
)

type AdminPage struct {
}

func (page *AdminPage) New(r *gin.Engine) {
	adminC := controllers.AdminController{}
	group := r.Group("/admin")
	{
		group.GET("/auth/login", adminC.Login)
	}
}

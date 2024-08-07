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
		group.POST("/auth/login", adminC.Login)
	}
}

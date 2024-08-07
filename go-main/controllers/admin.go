package controllers

import (
	"github.com/gin-gonic/gin"
	"rexai.com/helpers"
)

type AdminController struct {
}

func (control *AdminController) Login(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	instance := helpers.GetAuthInstance()
	if name != instance.GetName() || password != instance.GetSecret() {
		helpers.Respond(c, 401, "", nil)
		return
	}

	c.SetCookie("admin", "true", 3600, "/", "", false, true)
	helpers.Respond(c, 200, "", nil)
}

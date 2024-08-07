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

	session := &helpers.SessionHelper{}
	session.New(c)
	session.Set("admin", "1")

	helpers.Respond(c, 200, "", nil)
}

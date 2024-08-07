package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"rexai.com/helpers"
)

type AdminController struct {
}

type LoginFormParams struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (control *AdminController) Login(c *gin.Context) {
	var payload LoginFormParams
	err := c.BindJSON(&payload)
	if err != nil {
		helpers.Respond(c, 400, "参数错误", nil)
		return
	}
	instance := helpers.GetAuthInstance()
	log.Println(payload.Name, payload.Password, instance)
	if payload.Name != instance.GetName() || payload.Password != instance.GetSecret() {
		helpers.Respond(c, 401, "密码和用户名不匹配", nil)
		return
	}

	session := &helpers.SessionHelper{}
	session.New(c)
	session.Set("admin", "1")

	helpers.Respond(c, 200, "", nil)
}

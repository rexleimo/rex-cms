package main

import (
	"github.com/gin-gonic/gin"
	"rexai.com/helpers"
	"rexai.com/routes"
)

func main() {
	r := gin.Default()
	helpers.GetDbInstance()
	helpers.GetAuthInstance()
	helpers.RegisterSession(r)
	routes.RegisterRoutes(r)
	r.Run()
}

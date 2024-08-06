package controllers

import (
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"rexai.com/helpers"
	"rexai.com/models"
	"rexai.com/services"
)

var server *services.Image

type ImageController struct {
}

func init() {
	server = &services.Image{}
}

func (page *ImageController) List(c *gin.Context) {
	lists, err := server.List()
	if err != nil {
		helpers.Respond(c, 500, "", err)
	}
	helpers.Respond(c, 200, lists, nil)
}

func (page *ImageController) Create(c *gin.Context) {
	var image models.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	if err := server.Create(&image); err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	helpers.Respond(c, 200, image, nil)
}

func (page *ImageController) Update(c *gin.Context) {
	var image models.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if _, err := server.Update(int(id), image); err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	helpers.Respond(c, 200, image, nil)
}

func (page *ImageController) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := server.Delete(int(id)); err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	helpers.Respond(c, 200, "", nil)
}

func (page *ImageController) Show(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	image, err := server.Get(int(id))
	if err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	helpers.Respond(c, 200, image, nil)
}

func (page *ImageController) Upload(c *gin.Context) {
	file, handle, _ := c.Request.FormFile("file")
	defer file.Close()
	uploader := helpers.Upload{}
	content, err := io.ReadAll(file)
	if err != nil {
		helpers.Respond(c, 500, "", err)
		return
	}
	rawUrl, err := uploader.UploadFile(content, handle.Filename)
	if err != nil {
		helpers.Respond(c, 500, "", err)
	}
	helpers.Respond(c, 200, rawUrl, nil)
}

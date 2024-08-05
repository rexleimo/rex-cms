package routes

import "github.com/gin-gonic/gin"

type ImagePage struct {
}

func (page *ImagePage) New(r *gin.Engine) {
	group := r.Group("/img")
	{
		group.GET("", page.GetImageList)
	}
}

func (page *ImagePage) GetImageList(c *gin.Context) {

}

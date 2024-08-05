package helpers

import "github.com/gin-gonic/gin"

type Response[T any] struct {
	Status int         `json:"status"`
	Data   T           `json:"data"`
	Error  interface{} `json:"error,omitempty"`
}

func Respond[T any](c *gin.Context, status int, data T, err error) {
	var res Response[T]
	res.Status = status
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Data = data
	}
	c.JSON(status, res)
}

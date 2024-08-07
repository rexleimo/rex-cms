package middlewares

import (
	"github.com/gin-gonic/gin"
	"rexai.com/helpers"
)

func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否是管理员
		session := helpers.SessionHelper{}
		session.New(c)
		admin := session.Get("admin").(string)
		if admin == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "没有权限",
			})
			return
		}
		c.Next()
	}
}

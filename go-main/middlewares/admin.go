package middlewares

import "github.com/gin-gonic/gin"

func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否是管理员
		admin, err := c.Cookie("admin")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "没有权限",
			})
			return
		}
		if admin == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "没有权限",
			})
			return
		}
		c.Next()
	}
}

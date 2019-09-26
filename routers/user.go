package routers

import "github.com/gin-gonic/gin"

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func InitUserRouter(web *gin.Engine) {
	user := web.Group("/user")
	{
		user.POST("get", func(context *gin.Context) {
			var json Login
			if context.BindJSON(&json) == nil {

			}
		})
	}
}

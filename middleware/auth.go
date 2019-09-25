package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tdytaylor/go-gin-web-all/utils"
)

// jwt 认证

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.GetHeader("Auth")
		switch utils.RequestMethod(context) {
		case "GET":
		case "POST":
		default:
		}

	}
}

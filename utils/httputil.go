package utils

import "github.com/gin-gonic/gin"

func RequestMethod(context *gin.Context) string {
	return context.Request.Method
}

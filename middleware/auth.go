package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if strings.HasPrefix(context.Request.URL.String(), "/ws") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/login") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/webUser/login") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/webUser/register") ||
			strings.HasPrefix(context.Request.URL.String(), "/upload/") {
			context.Next()
			return
		}
		claims, err := common.GetCurrentUserClaims(context)
		if err != nil {
			context.AbortWithStatusJSON(200, gin.H{
				"code":    401,
				"message": "登陆过期，请重新登陆",
			})
			return
		}
		context.Set("claims", claims)
		context.Next()
	}
}

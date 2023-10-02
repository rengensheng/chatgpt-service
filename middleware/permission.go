package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/utils"
)

func RolePermission(roleValues []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		userClaims, exists := context.Get("claims")
		if exists {
			// 判断是否需要权限控制
			if len(roleValues) > 0 {
				hasRoleValues := userClaims.(*common.CustomClaims).Roles
				if utils.HasSameKey(roleValues, hasRoleValues) {
					context.Next()
					return
				} else {
					context.AbortWithStatusJSON(200, gin.H{
						"code":    403,
						"message": "无权限",
					})
					return
				}
			}
		}
		context.Next()
	}
}

func CodePermission(permissions []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		userClaims, exists := context.Get("claims")
		if exists {
			// 判断是否需要权限控制
			if len(permissions) > 0 {
				hasPermissions := userClaims.(*common.CustomClaims).Permission
				if utils.HasSameKey(permissions, hasPermissions) {
					context.Next()
					return
				} else {
					context.AbortWithStatusJSON(200, gin.H{
						"code":    403,
						"message": "无权限",
					})
					return
				}
			}
		}
		context.Next()
	}
}

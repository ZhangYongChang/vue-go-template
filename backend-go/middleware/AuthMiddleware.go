package middleware

import (
	"net/http"
	"strings"

	"github.com/ZhangYongChang/backend-go/common"
	"github.com/ZhangYongChang/backend-go/model"
	"github.com/ZhangYongChang/backend-go/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			util.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			util.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)
		if user.ID == 0 {
			util.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

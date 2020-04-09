package controller

import (
	"net/http"

	"github.com/ZhangYongChang/backend-go/dto"
	"github.com/ZhangYongChang/backend-go/model"
	"github.com/gin-gonic/gin"
)

func QueryUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

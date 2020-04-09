package controller

import (
	"log"
	"net/http"

	"github.com/ZhangYongChang/backend-go/common"
	"github.com/ZhangYongChang/backend-go/model"
	"github.com/ZhangYongChang/backend-go/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	user := model.User{}
	ctx.Bind(&user)
	name := user.Name
	telephone := user.Telephone
	password := user.Password

	if len(telephone) != 11 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须大于6位")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(db, telephone) {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		util.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}
	util.Success(ctx, gin.H{"token": token}, "注册成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

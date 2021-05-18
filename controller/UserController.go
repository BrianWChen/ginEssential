package controller

import (
	"brianwchen/ginessential/common"
	"brianwchen/ginessential/model"
	"brianwchen/ginessential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "telephone == 11"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "password >= 6"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "user existing"})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	log.Println(name, telephone, password)

	ctx.JSON(200, gin.H{
		"message": "successfully",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

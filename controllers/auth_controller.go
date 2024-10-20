package controllers

import (
	"exchangeapp/global"
	"exchangeapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "exchangeapp/models"

func Register(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPwd, err := utils.HashPassword(user.PassWord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.PassWord = hashPwd
	jwt, err := utils.GenerateJWT(user.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 完成 加密，token 生成
	if err := global.DB.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将用户插入数据库
	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": jwt})
	return
}

func Login(ctx *gin.Context) {
	// 查询用户名是否存在
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库中是否存在 用户名
	var user models.User
	// 这里的表名中的字段变成了 user_name 从查询的 sql 看出来的
	if err := global.DB.Where("user_name = ?", input.Username).First(&user).Error; err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名不存在"})
		return
	}
	// 验证密码
	if !utils.CheckPassword(user.PassWord, input.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名不存在"})
		return
	}
	// 生成 jwt token
	jwt, err := utils.GenerateJWT(user.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt})
	return
}

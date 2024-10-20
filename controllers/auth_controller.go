package controllers

import (
	"exchangeapp/utils"
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
}
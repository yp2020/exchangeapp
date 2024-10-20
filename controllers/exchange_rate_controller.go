package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateExchangeRate 创建汇率信息
func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	// 自动迁移，建立表
	if err := global.DB.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将数据插入数据库
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}

// GetExchangeRate 获取汇率数据
func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate
	if err := global.DB.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}

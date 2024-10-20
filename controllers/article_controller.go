package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 自动迁移，建立表
	if err := global.DB.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 插入数据库中
	if err := global.DB.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, article)
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	if err := global.DB.Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	var article models.Article
	if err := global.DB.Where("id = ?", id).First(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, article)
}

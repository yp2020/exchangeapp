package controllers

import (
	"exchangeapp/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// LikeArticle  为某个文章点赞
func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")
	likeKey := "article:" + articleId + ":likes"
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"result": "Successfully liked the article"})
}

// GetLikeArticleCount 获取某个文章的点赞数
func GetLikeArticleCount(ctx *gin.Context) {
	articleId := ctx.Param("id")
	likeKey := "article:" + articleId + ":likes"
	count, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		// Nil reply Redis returns when key does not exist.
		count = "0"
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"result": count})
}

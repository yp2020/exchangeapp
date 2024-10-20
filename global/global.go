package global

import "gorm.io/gorm"
import "github.com/go-redis/redis"

var (
	DB      *gorm.DB
	RedisDB *redis.Client
)

package config

import (
	"exchangeapp/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initDB(Host, Port, User, Password, DBName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User, Password, Host, Port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
		return
	}
	s, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
		return
	}
	s.SetMaxIdleConns(GlobalConfig.Database.MaxIdleConns)
	s.SetMaxOpenConns(GlobalConfig.Database.MaxOpenConns)
	s.SetConnMaxLifetime(time.Second * 60)
	global.DB = db
}

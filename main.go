package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
	"log"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		// 初始化失败，直接panic
		panic(err)
	}
	r := router.SetupRouter()
	port := config.GlobalConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(config.GlobalConfig.App.Port)
}

func testConfig() {
	log.Println(config.GlobalConfig.App.Name)
	log.Println(config.GlobalConfig.App.Port)

	log.Println(config.GlobalConfig.Database.Host)
	log.Println(config.GlobalConfig.Database.Port)
	log.Println(config.GlobalConfig.Database.User)
	log.Println(config.GlobalConfig.Database.Pass)
	log.Println(config.GlobalConfig.Database.Name)
}

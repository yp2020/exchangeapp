package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host         string
		Port         string
		User         string
		Pass         string
		Name         string
		MaxIdleConns int
		MaxOpenConns int
	}
	RedisConfig struct {
		Addr     string `yaml:"addr"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	}
}

// GlobalConfig 配置
var GlobalConfig = &Config{}

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalln("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatalf("config file found but another error was produced: %v", err)
		}
		return err
	}

	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalln("unmarshal config error")
		return err
	}
	// 初始化数据库
	initDB(GlobalConfig.Database.Host, GlobalConfig.Database.Port, GlobalConfig.Database.User, GlobalConfig.Database.Pass, GlobalConfig.Database.Name)
	InitRedis()
	return nil
}

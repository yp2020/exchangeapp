package main

import (
	"context"
	"exchangeapp/config"
	"exchangeapp/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		// 初始化失败，直接panic
		panic(err)
	}

	//testOrm()
	r := router.SetupRouter()
	port := config.GlobalConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	//r.Run(config.GlobalConfig.App.Port)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

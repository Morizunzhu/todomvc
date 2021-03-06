package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todomvc/client/core/router"
	"todomvc/core/dao"
)

func main() {
	router := router.InitRouter()

	dao.InitRedis()

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

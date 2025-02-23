package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var text = pflag.String("text", "Hello world!", "text to put on the webpage (optional) ")
var addr = pflag.String("addr", ":8080", "address to listen on (default :8080)")

func main() {
	pflag.Parse()
	if *text == "" {
		log.Fatal("--text option is required param!")
	}

	r := gin.Default()
	r.GET("/", TextHandler)
	r.GET("/health", HealthHandler)
	r.NoMethod(TextHandler)
	r.NoRoute(TextHandler)

	srv := http.Server{
		Addr:    *addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen failed: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server. Goodbye!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server exiting. Goodbye!!")
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OKAY!!",
		"version": "1.0.0",
		"timestamp": time.Now().Unix(),
		"uptime": time.Since(time.Now()),
		"message": "I'm alive!",
	})
}

func TextHandler(c *gin.Context) {
	c.String(http.StatusOK, *text)
}



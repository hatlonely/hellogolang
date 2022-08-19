package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Println("pre handler 1")
	})
	router.Use(func(c *gin.Context) {
		fmt.Println("pre handler 2")
	})
	router.Use(func(c *gin.Context) {
		c.Next()
		if c.IsAborted() {
			return
		}

		fmt.Println("post handler 1")
	})
	router.Use(func(c *gin.Context) {
		c.Next()
		if c.IsAborted() {
			return
		}

		fmt.Println("post handler 2")

		err := c.Errors.Last()
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "%v", c.Errors.Last())
			c.Abort()
		}
	})
	router.POST("/", func(c *gin.Context) {
		fmt.Println("do handler")
		c.Error(fmt.Errorf("error1"))
		c.Error(fmt.Errorf("error2"))
		c.Request.ParseForm()
		fmt.Println(c.Request.Form)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

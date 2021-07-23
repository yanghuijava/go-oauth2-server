package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/db"
	"go-oauth2-server/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Use(panicMiddleware)
	router.Use(loginMiddleware)

	web.NewHelloRoute().RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer db.Close()
	logrus.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}

func panicMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, err)
		}
	}()
	c.Next()
}

func loginMiddleware(c *gin.Context) {
	fmt.Println("执行前")
	c.Next()
	fmt.Println("执行后")
}

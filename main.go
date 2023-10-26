package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadConfig() {
	configFilePath := flag.String("config", "conf/", "Path to config file")
	flag.Parse()
	fmt.Println("Config file path:", *configFilePath)

	viper.SetConfigName("app")
	viper.AddConfigPath(*configFilePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	logrus.Info("Config file loaded successfully")
}

func startServer(engine *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	go func() {
		logrus.Info("Starting server on port ", port)
		// logrus.Info("Server host ", viper.GetString("server.host"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutting down server...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exiting")
}

func init() {
	loadConfig()
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())

	logrus.Info("Server context path ", viper.GetString("server.contextPath"))
	routes.SetupRoutes(engine)

	startServer(engine, viper.GetString("server.port"))
}

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"rentbook/internal/config"
	"rentbook/internal/factory"
	"rentbook/internal/middleware"
	"rentbook/utils/database/mysql"
	"rentbook/utils/pkg/redis"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Init() *config.AppConfig {
	// config
	getConfig := config.GetConfig()

	// logrus
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	var level logrus.Level = logrus.InfoLevel
	if config.GetConfig().LOGRUS_LEVEL != 0 {
		switch config.GetConfig().LOGRUS_LEVEL {
		case 1:
			level = 1
		case 2:
			level = 2
		case 3:
			level = 3
		case 4:
			level = 4
		case 5:
			level = 5
		case 6:
			level = 6
		}
	}
	logrus.SetLevel(level)

	return getConfig
}

func main() {
	// init
	cfg := Init()

	// db
	db := mysql.InitDB(cfg)

	// redis
	redis := redis.InitRedis(cfg)

	// echo
	e := echo.New()

	// factory
	factory.InitFactory(e, db, redis)

	// middleware
	middleware.Init(e)

	// start
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := e.Start(":" + strconv.Itoa(int(cfg.SERVER_PORT)))
		if err != nil {
			if err != http.ErrServerClosed {
				logrus.Fatal(err)
			}
		}
	}()

	<-ch

	logrus.Println("Shutting down server...")
	cancel()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	e.Shutdown(ctx2)
	logrus.Println("Server gracefully stopped")
}

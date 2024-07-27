package mysql

import (
	"fmt"
	"rentbook/internal/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	var level logger.LogLevel = 4
	if config.GetConfig().GORM_LEVEL != 0 {
		switch config.GetConfig().GORM_LEVEL {
		case 1:
			level = 1
		case 2:
			level = 2
		case 3:
			level = 3
		case 4:
			level = 4
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database %s", config.DB_DRIVER))
	}
	logrus.Info(fmt.Sprintf("Successfully connected to database %s", config.DB_DRIVER))

	return db
}

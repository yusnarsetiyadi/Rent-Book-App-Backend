package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	GORM_LEVEL     int
	LOGRUS_LEVEL   int
	APP            string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_HOST        string
	DB_PORT        int
	DB_NAME        string
	SERVER_PORT    int
	JWT_SECRET     string
	REDIS_ADDRESS  string
	REDIS_PORT     int
	REDIS_USER     string
	REDIS_PASSWORD string
	SMTP_HOST      string
	SMTP_PORT      uint
	SENDER_NAME    string
	AUTH_EMAIL     string
	AUTH_PASSWORD  string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	if _, exist := os.LookupEnv("SECRET"); !exist {
		if err := godotenv.Load(".env"); err != nil {
			log.Println(err)
		}
	}

	cnvGormLevel, err := strconv.Atoi(os.Getenv("GORM_LEVEL"))
	if err != nil {
		log.Fatal("Cannot parse Gorm Level variable")
		return nil
	}
	cnvLogrusLevel, err := strconv.Atoi(os.Getenv("LOGRUS_LEVEL"))
	if err != nil {
		log.Fatal("Cannot parse Logrus Level variable")
		return nil
	}
	defaultConfig.GORM_LEVEL = cnvGormLevel
	defaultConfig.LOGRUS_LEVEL = cnvLogrusLevel
	defaultConfig.APP = os.Getenv("APP")
	cnvServerPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Cannot parse Server Port variable")
		return nil
	}
	defaultConfig.SERVER_PORT = cnvServerPort
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	cnvDBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.DB_PORT = cnvDBPort
	defaultConfig.JWT_SECRET = os.Getenv("JWT_SECRET")
	defaultConfig.REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
	cnvRedisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatal("Cannot parse Redis Port variable")
		return nil
	}
	defaultConfig.REDIS_PORT = cnvRedisPort
	defaultConfig.REDIS_USER = os.Getenv("REDIS_USER")
	defaultConfig.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	defaultConfig.SMTP_HOST = os.Getenv("SMTP_HOST")
	cnvSmtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Cannot parse SMTP Port variable")
		return nil
	}
	defaultConfig.SMTP_PORT = uint(cnvSmtpPort)
	defaultConfig.SENDER_NAME = os.Getenv("SENDER_NAME")
	defaultConfig.AUTH_EMAIL = os.Getenv("AUTH_EMAIL")
	defaultConfig.AUTH_PASSWORD = os.Getenv("AUTH_PASSWORD")
	return &defaultConfig
}

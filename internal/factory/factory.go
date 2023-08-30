package factory

import (
	userDelivery "rentbook/internal/features/user/delivery"
	userRepo "rentbook/internal/features/user/repository"
	userService "rentbook/internal/features/user/service"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, redis *redis.Client) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory, db, redis)
	userDelivery.New(userServiceFactory, e)
}

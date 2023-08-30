package factory

import (
	userDelivery "rentbook/internal/features/user/delivery"
	userRepo "rentbook/internal/features/user/repository"
	userService "rentbook/internal/features/user/service"

	bookDelivery "rentbook/internal/features/book/delivery"
	bookRepo "rentbook/internal/features/book/repository"
	bookService "rentbook/internal/features/book/service"

	authDelivery "rentbook/internal/features/auth/delivery"
	authService "rentbook/internal/features/auth/service"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, redis *redis.Client) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory, redis)
	userDelivery.New(userServiceFactory, e)

	bookRepoFactory := bookRepo.New(db)
	bookServiceFactory := bookService.New(bookRepoFactory, redis)
	bookDelivery.New(bookServiceFactory, e)

	authServiceFactory := authService.New(userRepoFactory, redis)
	authDelivery.New(authServiceFactory, e)
}

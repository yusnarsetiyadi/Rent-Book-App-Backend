package factory

import (
	userDelivery "rentbook/internal/features/user/delivery"
	userRepo "rentbook/internal/features/user/repository"
	userService "rentbook/internal/features/user/service"

	bookDelivery "rentbook/internal/features/book/delivery"
	bookRepo "rentbook/internal/features/book/repository"
	bookService "rentbook/internal/features/book/service"

	rentDelivery "rentbook/internal/features/rent/delivery"
	rentRepo "rentbook/internal/features/rent/repository"
	rentService "rentbook/internal/features/rent/service"

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

	rentRepoFactory := rentRepo.New(db)
	rentServiceFactory := rentService.New(rentRepoFactory, redis)
	rentDelivery.New(rentServiceFactory, e)

	authServiceFactory := authService.New(userRepoFactory, rentRepoFactory, redis)
	authDelivery.New(authServiceFactory, e)
}

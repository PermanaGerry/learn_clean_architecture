package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"learn-hexagonal-architecture/internal/delivery/http"
	"learn-hexagonal-architecture/internal/delivery/http/route"
	"learn-hexagonal-architecture/internal/repository"
	"learn-hexagonal-architecture/internal/usecase"
)

type BootstrapConfig struct {
	Config    *viper.Viper
	App       *fiber.App
	DB        *gorm.DB
	Log       *logrus.Logger
	Validator *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// setup repository
	userRepository := repository.NewUserRepository(config.Log)

	//	setup usecase
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validator, userRepository)

	//	setup controller
	userController := http.NewUserController(config.Log, userUseCase)

	//	 setup route
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}

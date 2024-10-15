package main

import (
	"fmt"
	"learn-hexagonal-architecture/config"
	"strconv"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		Config:    viperConfig,
		App:       app,
		DB:        db,
		Log:       log,
		Validator: validator,
	})

	webPort, err := strconv.Atoi(viperConfig.GetString("APP_PORT"))
	if err != nil {
		log.Fatalf("Failed to conncet databse: %v", err)
	}

	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

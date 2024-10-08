package main

import (
	"learn-hexagonal-architecture/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger()
	db := config.NewDatabase(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		Config: viperConfig,
		DB:     db,
	})

}

package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Config *viper.Viper
	DB     *gorm.DB
	Log    *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	fmt.Print(config)
}

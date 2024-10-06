package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

var poolIdle, poolMaxConn, poolLifetime int

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := viper.GetString("MYSQL_HOST")
	port := viper.GetString("MYSQL_PORT")
	username := viper.GetString("MYSQL_USERNAME")
	password := viper.GetString("MYSQL_PASSWORD")
	database := viper.GetString("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to conncet databse: %v", err)
	}

	if poolIdle, err = strconv.Atoi(viper.GetString("MYSQL_POOL_IDLE")); err != nil {
		log.Fatal("Format wrong mysql pool idle")
	}

	if poolMaxConn, err = strconv.Atoi(viper.GetString("MYSQL_POOL_MAX_CONNECTION")); err != nil {
		log.Fatal("Format wrong mysql pool max connection")
	}

	if poolLifetime, err = strconv.Atoi(viper.GetString("MYSQL_POOL_LIFETIME")); err != nil {
		log.Fatal("Format wrong mysql pool lifetime")
	}

	connection.SetMaxIdleConns(poolIdle)
	connection.SetMaxOpenConns(poolMaxConn)
	connection.SetConnMaxIdleTime(time.Second * time.Duration(poolLifetime))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args)
}

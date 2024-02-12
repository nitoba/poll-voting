package configs

import (
	"context"

	"github.com/spf13/viper"
)

var (
	config *conf
	logger *Logger
)

type conf struct {
	// DB configuration
	DATABASE_URL string
	// Web server configuration
	Port string `mapstructure:"PORT"`

	// JWT configuration
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	Ctx          context.Context
}

func LoadConfig(path ...string) (*conf, error) {
	logger := GetLogger("configs")
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	if len(path) == 0 {
		logger.Warn("Running in test environment...")
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(path[0])
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Error to reading configs: %v", err)
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Errorf("Error to reading configs: %v", err)
		panic(err)
	}
	config.Ctx = context.Background()
	return config, err
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetConfig() *conf {
	return config
}

package configs

import (
	"context"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	config *conf
	logger *Logger
)

type conf struct {
	ENV string `mapstructure:"ENV"`
	// DB configuration
	DATABASE_URL string

	// Redis configuration
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDb       int    `mapstructure:"REDIS_DB"`

	// Web server configuration
	Port string `mapstructure:"PORT"`

	// JWT configuration
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	Ctx          context.Context
}

func LoadConfig(envPath ...string) (*conf, error) {
	logger := GetLogger("configs")
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	rootDir := rootDir()
	if len(envPath) == 0 {
		viper.SetConfigFile(rootDir + "/.env")
	} else {
		viper.SetConfigFile(rootDir + "/" + envPath[0])
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

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetConfig() *conf {
	return config
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfigurations struct {
	Server      Server      `mapstructure:"server"`
	RateLimiter RateLimiter `mapstructure:"rateLimiter"`
}

type Server struct {
	Host string `mapstructure:"host"`
}

type RateLimiter struct {
	Strategy    string      `mapstructure:"strategy"`
	TokenBucket TokenBucket `mapstructure:"tokenBucket"`
}

type TokenBucket struct {
	ReqPerMin int `mapstructure:"reqPerMin"`
}

var AppConfig AppConfigurations

func LoadAppConfigurations(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config:%v", err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("failed to unmarshal config:%v", err)
	}
}

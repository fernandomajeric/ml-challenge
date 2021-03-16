package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	GeneralConfig struct {
		App        App       `mapstructure: "app"`
		RedisCache RedisRing `mapstructure:"redis_cache"`
		SecretKey  string    `mapstructure:"secret_key"`
	}
)

type App struct {
	Env  string
	Rest Rest
}

type Rest struct {
	Country         Country
	Currency        Currency
	Geolocalization Geolocalization
}
type Country struct {
	Url string
}

type Currency struct {
	Url   string
	Token string
}
type Geolocalization struct {
	Url string
}

var Configuration GeneralConfig

func LoadConfig(filepaths ...string) *GeneralConfig {
	if len(filepaths) == 0 {
		panic(fmt.Errorf("Empty config file"))
	}
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(filepaths[0])

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}
	err := v.Unmarshal(&Configuration)
	if err != nil {
		panic(fmt.Errorf("Fatal error marshal config file: %s \n", err))
	}
	return &Configuration
}

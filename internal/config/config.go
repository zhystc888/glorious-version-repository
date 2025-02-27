package config

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func Load() *Config {
	var configFile string
	flag.StringVar(&configFile, "c", "config/config.yaml", "config file")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		Port: v.GetString("port"),
	}
}

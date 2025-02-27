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
	flag.StringVar(&configFile, "c", "config.yaml", "config file")
	flag.Parse()

	v := viper.New()
	v.SetConfigName(configFile)
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		Port: v.GetString("port"),
	}
}

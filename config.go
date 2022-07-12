package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *Config

func LoadConfig() {
	log.Println("Memuat Server Config...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}

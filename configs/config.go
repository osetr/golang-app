package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBAddr string
	DBUser string
	DBPass string
	DBName string
}

func NewConfig(fileName string) *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName(fileName)
	viper.Set("db.password", os.Getenv("password"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	return &Config{
		DBAddr: viper.GetString("db.addr"),
		DBUser: viper.GetString("db.user"),
		DBName: viper.GetString("db.database"),
		DBPass: os.Getenv("password"),
	}
}

package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfig(fileName string) {
	viper.AddConfigPath("configs")
	viper.SetConfigName(fileName)
	viper.Set("db.password", os.Getenv("password"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}

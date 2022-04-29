package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func HydrateFromConfig[T any](configPath string, config *T) {

	viper.AddConfigPath("./")
	viper.SetConfigType("yml")
	viper.SetConfigName(configPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	viper.Unmarshal(&config)
}

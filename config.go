package ipage

import (
	"log"

	"github.com/spf13/viper"
)

// Getconfig return a map of configuration settings from a configuration file
func Getconfig(cfg string) map[string]any {
	viper.SetConfigFile(cfg)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	msite := viper.AllSettings()
	return msite

}

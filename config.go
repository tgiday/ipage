package magpie

import (
	"log"

	"github.com/spf13/viper"
)

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

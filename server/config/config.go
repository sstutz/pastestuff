package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/sstutz/pastestuff/server/helpers"
)

type config struct {
	Server struct {
		Address string
		Port    int
	}
}

var Config config

func init() {
	viper.SetConfigName("testconfig")
	setConfigPaths()
	loadConfiguration()

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func setConfigPaths() {
	viper.AddConfigPath("/etc/" + helpers.AppName + "/")
	viper.AddConfigPath("$HOME/.config/")
}

func loadConfiguration() {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		setDefaultConfiguration()
	}
}

func setDefaultConfiguration() {
	viper.SetDefault("server.address", "localhost")
	viper.SetDefault("server.port", "3003")
}

package helpers

import (
	"github.com/Andrem19/Timer_Job/variables"
	"github.com/spf13/viper"
)


func LoadConfig(path string) (*variables.Conf, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		AddToLog(err.Error())
	}

	err = viper.Unmarshal(&variables.Config)
	return &variables.Config, nil
}
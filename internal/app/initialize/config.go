package initialize

import (
	"fmt"

	"github.com/new-pop-corn/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./config/%s-debug.yaml", configFilePrefix)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic(err)
	}

	if err := v.Unmarshal(&config.ServerConf); err != nil {
		zap.S().Panic(err)
	}
}

package config_base

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

const _CONFIG_PATH = "./config.toml"

type config struct {
	V *viper.Viper
}

var (
	once     sync.Once
	instance *config
)

func Instance() *config {
	once.Do(func() {
		// check config file "./config.toml"
		_, err := os.Stat(_CONFIG_PATH)
		if err != nil {
			if os.IsNotExist(err) {
				_, err = os.Create(_CONFIG_PATH)
				if err != nil {
					panic(fmt.Errorf("config initialization failed, try to create config file with error: %v", err))
				}
			} else {
				panic(fmt.Errorf("config initialization failed with error: %v", err))
			}
		}

		// viper init
		v := viper.New()
		v.SetConfigFile(_CONFIG_PATH)

		err = v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("config initialization failed with error: %v", err))
		}

		// config init
		instance = &config{
			V: v,
		}
	})

	return instance
}

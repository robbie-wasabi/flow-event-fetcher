package config

import (
	"fmt"
	"unicode/utf8"

	"github.com/spf13/viper"
	// _ "github.com/spf13/viper/remote"
)

const (
	PATH_FLOW_JSON = "../flow.json"
)

var Conf *Config

func InitConf(env string) {
	Conf = new(Config)
	Conf.setEnv(env)
	if err := Conf.setupViper(env); err != nil {
		fmt.Println(err)
	}
}

type Config struct {
	Env            string
	FlowAccessNode string
}

func (c *Config) setEnv(env string) {
	c.Env = env
}

func (c *Config) setupViper(env string) error {
	fmt.Println(fmt.Sprintf("Environment: %s", env))

	viper.SetConfigFile(PATH_FLOW_JSON)
	if err := viper.ReadInConfig(); err != nil {
		if err.Error() != fmt.Sprintf("open %v: no such file or directory", PATH_FLOW_JSON) {
			panic(err)
		}
		// handle filepath for cli
		viper.SetConfigFile(trimFirstRune(PATH_FLOW_JSON))
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	// env
	c.FlowAccessNode = viper.GetString(fmt.Sprintf("networks.%s", env))

	// todo: cleanup
	if env == "testnet" {
	} else {
	}

	return nil
}

func (c *Config) GetEnv() string {
	return c.Env
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

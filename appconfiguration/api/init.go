package api

import "github.com/spf13/viper"

func init() {
	viper.SetDefault(PortKey, DefaultPort)
}

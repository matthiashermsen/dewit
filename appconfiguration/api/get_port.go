package api

import "github.com/spf13/viper"

func GetPort() uint16 {
	port := viper.GetUint16(PortKey)

	return port
}

package logging

import "github.com/spf13/viper"

func init() {
	viper.SetDefault(LevelKey, DefaultLevel)
}

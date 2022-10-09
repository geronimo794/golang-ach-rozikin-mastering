package utils

import "github.com/spf13/viper"

func GetValue(key string) string {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return viper.GetString(key)
}

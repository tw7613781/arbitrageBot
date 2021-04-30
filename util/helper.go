package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL string `json:"BaseURL"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Fail to unmarshal json to struct %s", err)
	}

	return config
}

func SortMapByKeyToString(params map[string]interface{}) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}

	output := ""
	for _, k := range keys {
		output += "\"" + k + "\":" + "\"" + params[k].(string) + "\""
	}

	return output
}

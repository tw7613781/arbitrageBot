package util

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL string `json:"BaseURL"`
}

func GetConfig(path string) *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatalf("Fail to unmarshal json to struct %s", err)
	}

	return c
}

func GetTimestampMili() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

package util

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	BaseURL string `json:"BaseURL"`
}

func GetConfig() *config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	c := &config{}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatalf("Fail to unmarshal json to struct %s", err)
	}

	return c
}

func GetTimestampMili() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

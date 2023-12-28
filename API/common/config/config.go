package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	ServerURL string `json:"server_url"`
	MongoURL  string `json:"mongo_url"`
}

var GlobalConfig Config

func InitGlobalConfig(configPath string) {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &GlobalConfig)
	if err != nil {
		panic(err)
	}
}

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitGlobalConfig(t *testing.T) {
	filePath := "../../conf/local_config.json"
	InitGlobalConfig(filePath)
	wantConfig := Config{
		ServerURL: "localhost:8100",
		MongoURL:  "mongodb://localhost:27017",
	}
	assert.Equal(t, wantConfig, GlobalConfig)
}

package config

import (
	"testing"
)

func TestConfigLoad(t *testing.T) {
	LoadConfig("./dev.yaml")
	t.Log(ServerConfig)
}

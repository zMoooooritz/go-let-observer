package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Config *Configuration
)

type Configuration struct {
	ServerCredentials ServerCredentials `yaml:"ServerCredentials"`
	UIStartupOptions  UIStartupOptions  `yaml:"UIStartupOptions"`
}

type ServerCredentials struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type UIStartupOptions struct {
	ScreenSize            int  `yaml:"ScreenSize"`
	ShowPlayers           bool `yaml:"ShowPlayers"`
	ShowPlayerInfo        bool `yaml:"ShowPlayerInfo"`
	ShowSpawns            bool `yaml:"ShowSpawns"`
	ShowGridOverlay       bool `yaml:"ShowGridOverlay"`
	ShowServerInfoOverlay bool `yaml:"ShowServerInfoOverlay"`
}

func InitConfig(configFile string) error {
	Config = defaultConfiguration()
	if configFile == "" {
		return nil
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("Configuration error: %s", err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return fmt.Errorf("Configuration error: %s", err)
	}

	return nil
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		UIStartupOptions: UIStartupOptions{
			ScreenSize:            1000,
			ShowPlayers:           true,
			ShowPlayerInfo:        true,
			ShowSpawns:            false,
			ShowGridOverlay:       true,
			ShowServerInfoOverlay: false,
		},
	}
}

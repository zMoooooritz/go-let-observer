package util

import (
	"fmt"
	"os"

	"github.com/zMoooooritz/go-let-loose/pkg/rconv2"
	"gopkg.in/yaml.v3"
)

var (
	Config *Configuration
)

type Configuration struct {
	ServerCredentials ServerCredentials `yaml:"ServerCredentials"`
	UIOptions         UIOptions         `yaml:"UIStartupOptions"`
	ReplaysDirectory  string            `yaml:"ReplaysDirectory"`
}

func (c *Configuration) GetServerCredentials() rconv2.ServerConfig {
	return rconv2.ServerConfig{
		Host:     c.ServerCredentials.Host,
		Port:     c.ServerCredentials.Port,
		Password: c.ServerCredentials.Password,
	}
}

type ServerCredentials struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type UIOptions struct {
	ScreenSize            int  `yaml:"ScreenSize"`
	ShowPlayers           bool `yaml:"ShowPlayers"`
	ShowPlayerInfo        bool `yaml:"ShowPlayerInfo"`
	ShowSpawns            bool `yaml:"ShowSpawns"`
	ShowTanks             bool `yaml:"ShowTanks"`
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
		UIOptions: UIOptions{
			ScreenSize:            1000,
			ShowPlayers:           true,
			ShowPlayerInfo:        true,
			ShowSpawns:            false,
			ShowTanks:             true,
			ShowGridOverlay:       true,
			ShowServerInfoOverlay: false,
		},
	}
}

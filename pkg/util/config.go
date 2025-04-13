package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/yaml.v3"
)

var (
	Config *Configuration
)

type Configuration struct {
	ServerCredentials ServerCredentials `yaml:"ServerCredentials"`
	UIStartupOptions  UIStartupOptions  `yaml:"UIStartupOptions"`
	Keys              Keys              `yaml:"Keys"`
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

type Keys struct {
	IncreaseInterval        []string `yaml:"IncreaseInterval"`
	DecreaseInterval        []string `yaml:"DecreaseInterval"`
	TogglePlayers           []string `yaml:"TogglePlayers"`
	TogglePlayerInfo        []string `yaml:"TogglePlayerInfo"`
	ToggleSpawns            []string `yaml:"ToggleSpawns"`
	ToggleGridOverlay       []string `yaml:"ToggleGridOverlay"`
	ToggleServerInfoOverlay []string `yaml:"ToggleServerInfoOverlay"`
	ShowScoreboard          []string `yaml:"ShowScoreboard"`
	Quit                    []string `yaml:"Quit"`
	Help                    []string `yaml:"Help"`
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
		Keys: defaultKeys(),
	}
}

func defaultKeys() Keys {
	return Keys{
		IncreaseInterval:        []string{"+"},
		DecreaseInterval:        []string{"-"},
		TogglePlayers:           []string{"p"},
		TogglePlayerInfo:        []string{"i"},
		ToggleSpawns:            []string{"s"},
		ToggleGridOverlay:       []string{"g"},
		ToggleServerInfoOverlay: []string{"h"},
		ShowScoreboard:          []string{"tab"},
		Quit:                    []string{"q"},
		Help:                    []string{"?"},
	}
}

func MapKey(key string) ebiten.Key {
	key = strings.ToLower(key)
	switch key {
	case "tab":
		return ebiten.KeyTab
	case "enter":
		return ebiten.KeyEnter
	case "escape", "esc":
		return ebiten.KeyEscape
	case "space":
		return ebiten.KeySpace
	default:
		if len(key) == 1 {
			return ebiten.Key(key[0])
		}
	}
	return ebiten.KeyMax
}

package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	Game struct {
		ScreenWidth  int64 `yaml:"SCREENWIDTH"`
		ScreenHeight int64 `yaml:"SCREENHEIGHT"`
	} `yaml:"GAME"`
	Player struct {
		Player1 string `yaml:"PLAYER1"`
	} `yaml:"PLAYERS"`
}

func init() {
	err := loadConfigFromYAML("./config.yaml")
	if err != nil {
		log.Fatal("Failed to load YAML: ", err)
	}

	fmt.Printf("Final Merged Config: %+v\n", cfg.Game.ScreenHeight)
}

func loadConfigFromYAML(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	return nil
}

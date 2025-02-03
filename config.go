package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ScreenWidth      int `json:"ScreenWidth"`
	ScreenHeight     int `json:"ScreenHeight"`
	GridSize         int `json:"GridSize"`
	GameSpeed        int `json:"GameSpeed"`
	StartSnakeLength int `json:"StartSnakeLength"`
}

var config Config

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Println("Loaded config data:", string(data))

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	fmt.Printf("Config loaded: %+v\n", config)

	if config.ScreenWidth <= 0 || config.ScreenHeight <= 0 {
		return fmt.Errorf("Invalid config values: ScreenWidth and ScreenHeight must be positive")
	}

	return nil
}

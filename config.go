package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"os"
)

type Config struct {
	ScreenWidth           int  `json:"ScreenWidth"`
	ScreenHeight          int  `json:"ScreenHeight"`
	GridSize              int  `json:"GridSize"`
	GameSpeed             int  `json:"GameSpeed"`
	StartSnakeLength      int  `json:"StartSnakeLength"`
	GameOverOnBorderTouch bool `json:"GameOverOnBorderTouch"`
	GameOverOnTailTouch   bool `json:"GameOverOnTailTouch"`
}

var (
	config     Config
	arcadeFont font.Face
)

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

	// Loading font
	fontBytes, err := os.ReadFile("assets/arcadeFont.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(config.ScreenWidth / 30),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

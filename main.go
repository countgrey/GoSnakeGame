package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

func main() {
	err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	game := &Game{
		snake:    Snake{body: []Point{{config.ScreenWidth / config.GridSize / 2, config.ScreenHeight / config.GridSize / 2}}, direction: Right, color: color.RGBA{255, 0, 0, 255}},
		food:     Point{config.ScreenWidth / config.GridSize / 3, config.ScreenHeight / config.GridSize / 3},
		gameOver: false,
	}

	for i := 0; i < config.StartSnakeLength; i++ {
		game.snake.move(true)
	}

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) Update() error {
	g.tick++
	if g.tick%config.GameSpeed != 0 {
		return nil
	}

	//Handle input
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.direction != Down {
		g.snake.direction = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.direction != Up {
		g.snake.direction = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.direction != Right {
		g.snake.direction = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.direction != Left {
		g.snake.direction = Right
	}

	if g.snake.body[0] != g.food {
		g.snake.move(false)
	} else {
		g.snake.move(true)
		g.food.X = rand.Intn(config.ScreenWidth) / config.GridSize
		g.food.Y = rand.Intn(config.ScreenHeight) / config.GridSize
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for _, point := range g.snake.body {
		ebitenutil.DrawRect(screen, float64(point.X*config.GridSize), float64(point.Y*config.GridSize), float64(config.GridSize), float64(config.GridSize), g.snake.color)
	}

	ebitenutil.DrawRect(screen, float64(g.food.X*config.GridSize), float64(g.food.Y*config.GridSize), float64(config.GridSize), float64(config.GridSize), color.RGBA{0, 255, 0, 255})
}

func (snake *Snake) move(expand bool) {
	head := snake.body[0]
	head.X += directionVectors[snake.direction].X
	head.Y += directionVectors[snake.direction].Y
	snake.body = append([]Point{head}, snake.body...)
	if !expand {
		snake.body = snake.body[:len(snake.body)-1]
	}
}

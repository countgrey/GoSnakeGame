package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"math/rand"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.restart()
		}
		return nil
	}

	// Game speed setting
	g.tick++
	if g.tick%config.GameSpeed != 0 {
		return nil
	}

	// Handle input
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.direction != Down {
		g.snake.direction = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.direction != Up {
		g.snake.direction = Down
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.direction != Right {
		g.snake.direction = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.direction != Left {
		g.snake.direction = Right
	}

	// Checking if snake move out of game border
	head := &g.snake.body[0]

	if config.GameOverOnBorderTouch &&
		(head.X < 0 || head.X >= config.ScreenWidth/config.GridSize ||
			head.Y < 0 || head.Y >= config.ScreenHeight/config.GridSize) {
		g.gameOver = true
		return nil
	}

	// Teleporting to other side of the screen
	head.X = (head.X + config.ScreenWidth/config.GridSize) % (config.ScreenWidth / config.GridSize)
	head.Y = (head.Y + config.ScreenHeight/config.GridSize) % (config.ScreenHeight / config.GridSize)

	// Checking if snake eats its own tail
	if config.GameOverOnTailTouch {
		for _, cell := range g.snake.body[1:] {
			if g.snake.body[0] == cell {
				g.gameOver = true
				return nil
			}
		}
	}

	// Food eating
	if g.snake.body[0] != g.food {
		g.snake.move(false)
	} else {
		g.score++
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

	if g.gameOver {
		text.Draw(screen, "GAME OVER", arcadeFont, config.ScreenWidth/2-40, config.ScreenHeight/2, color.RGBA{255, 0, 0, 255})
		text.Draw(screen, fmt.Sprintf("SCORE: %d", g.score), arcadeFont, config.ScreenWidth/2-40, config.ScreenHeight/2+20, color.White)
		text.Draw(screen, "PRESS ENTER TO RESTART", arcadeFont, config.ScreenWidth/2-40, config.ScreenHeight/2+40, color.White)
	}

	text.Draw(screen, fmt.Sprintf("SCORE: %d", g.score), arcadeFont, 20, 20, color.White)
}

func (g *Game) restart() {
	g.gameOver = false
	g.score = 0

	g.snake = Snake{
		body:      []Point{{config.ScreenWidth / config.GridSize / 2, config.ScreenHeight / config.GridSize / 2}},
		direction: Right,
		color:     color.RGBA{255, 0, 0, 255},
	}

	g.food = Point{
		X: rand.Intn(config.ScreenWidth / config.GridSize),
		Y: rand.Intn(config.ScreenHeight / config.GridSize),
	}

	for i := 0; i < config.StartSnakeLength; i++ {
		g.snake.move(true)
	}
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

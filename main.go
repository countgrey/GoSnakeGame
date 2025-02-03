package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
	GameSpeed    = 5
)

type Point struct {
	X, Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var directionVectors = map[Direction]Point{
	Up:    {0, -1},
	Down:  {0, 1},
	Left:  {-1, 0},
	Right: {1, 0},
}

type Snake struct {
	body      []Point
	direction Direction
	color     color.Color
}

type Game struct {
	snake    Snake
	food     Point
	gameOver bool
	tick     int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.tick++
	if g.tick%GameSpeed != 0 {
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

	head := g.snake.body[0]
	head.X += directionVectors[g.snake.direction].X
	head.Y += directionVectors[g.snake.direction].Y
	g.snake.body = append([]Point{head}, g.snake.body...)

	if head != g.food {
		g.snake.body = g.snake.body[:len(g.snake.body)-1]
	} else {
		g.food.X = rand.Intn(ScreenWidth) / GridSize
		g.food.Y = rand.Intn(ScreenHeight) / GridSize
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for _, point := range g.snake.body {
		ebitenutil.DrawRect(screen, float64(point.X*GridSize), float64(point.Y*GridSize), GridSize, GridSize, g.snake.color)
	}

	ebitenutil.DrawRect(screen, float64(g.food.X*GridSize), float64(g.food.Y*GridSize), GridSize, GridSize, color.RGBA{0, 255, 0, 255})
}

func main() {
	game := &Game{
		snake:    Snake{body: []Point{{ScreenWidth / GridSize / 2, ScreenHeight / GridSize / 2}}, direction: Right, color: color.RGBA{255, 0, 0, 255}},
		food:     Point{ScreenWidth / GridSize / 3, ScreenHeight / GridSize / 3},
		gameOver: false,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import "image/color"

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
	score    int
}

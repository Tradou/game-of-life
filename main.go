package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	rows = 50
	cols = 100
	size = 10
)

type Grid [][]bool

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, float64(cols*size), float64(rows*size)),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := generateCell()

	for !win.Closed() {
		draw(win, grid)
	}
}

func generateCell() Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) == 0
		}
	}
	return grid
}

func draw(win *pixelgl.Window, grid Grid) {
	win.Clear(pixel.RGB(1, 1, 1)) // Blanc

	for i, row := range grid {
		for j, alive := range row {
			if alive {
				drawCell(win, j*size, (rows-i-1)*size)
			}
		}
	}

	win.Update()
}

func drawCell(win *pixelgl.Window, x, y int) {
	rect := pixel.R(float64(x), float64(y), float64(x+size), float64(y+size))
	sprite := pixel.NewSprite(nil, rect)
	sprite.Draw(win, pixel.IM.Moved(rect.Center()).Scaled(rect.Center(), 0.5))
}

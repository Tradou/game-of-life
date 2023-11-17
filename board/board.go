package board

import (
	"game-of-life/mutation"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

const pGenerate = 0.5

type Board struct {
	Grid  Grid
	Rules Ruler
}

type Grid [][]mutation.Cell

type Ruler interface {
	UnderPopulation(grid Grid, i, j int) bool
	OverPopulation(grid Grid, i, j int) bool
	Reproduce(grid Grid, i, j int) bool
}

const (
	Rows = 50
	Cols = 100
	Size = 10
)

func New(r Ruler) *Board {
	return &Board{
		Grid:  GenerateCell(),
		Rules: r,
	}
}

func GenerateCell() Grid {
	grid := make(Grid, Rows)
	for i := range grid {
		grid[i] = make([]mutation.Cell, Cols)
		for j := range grid[i] {
			if rand.Intn(100) <= pGenerate*100 {
				grid[i][j].State = "ALIVE"
				if mutation.CanMutate(grid[i][j]) {
					mutation.FindMutation(&grid[i][j])
				}
			} else {
				grid[i][j].State = "DEAD"
			}
		}
	}
	return grid
}

func (b *Board) Draw(win *pixelgl.Window) {
	win.Clear(pixel.RGB(1, 1, 1))

	for i, row := range b.Grid {
		for j := range row {
			if isAlive(b.Grid, i, j) {
				b.DrawCell(win, j*Size, (Rows-i-1)*Size)
			}
		}
	}

	win.Update()
}

func (b *Board) DrawCell(win *pixelgl.Window, x, y int) {
	rect := pixel.R(float64(x), float64(y), float64(x+Size), float64(y+Size))
	sprite := pixel.NewSprite(nil, rect)
	sprite.Draw(win, pixel.IM.Moved(rect.Center()).Scaled(rect.Center(), 0.5))
}

func (b *Board) Update() {
	newGrid := make(Grid, Rows)
	for i := range b.Grid {
		newGrid[i] = make([]mutation.Cell, Cols)
		copy(newGrid[i], b.Grid[i])

		for j := range b.Grid[i] {
			if isAlive(b.Grid, i, j) {
				if b.Rules.UnderPopulation(b.Grid, i, j) || b.Rules.OverPopulation(b.Grid, i, j) {
					newGrid[i][j].State = "DEAD"
				}
			} else {
				if b.Rules.Reproduce(b.Grid, i, j) {
					newGrid[i][j].State = "ALIVE"
					if mutation.CanMutate(b.Grid[i][j]) {
						mutation.FindMutation(&newGrid[i][j])
					}
				}
			}
		}
	}

	b.Grid = newGrid
}

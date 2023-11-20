package board

import (
	"game-of-life/mutation"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
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
	Reproduce(grid Grid, i, j int) (bool, int, int)
	DieFromInstability(grid Grid, i, j int) bool
	WarriorInvasion() bool
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
				create(&grid[i][j])
				if mutation.CanMutate(grid[i][j], 0) {
					mutation.FindMutation(&grid[i][j], 0)
				}
			} else {
				kill(&grid[i][j])
			}
		}
	}
	return grid
}

func (b *Board) Draw(win *pixelgl.Window) {
	win.Clear(pixel.RGB(1, 1, 1))

	for i, row := range b.Grid {
		for j := range row {
			if isMutant(b.Grid[i][j]) {
				b.DrawCell(win, j*Size, (Rows-i-1)*Size, color.RGBA{G: 255, A: 255})
			} else if isAlive(b.Grid[i][j]) {
				b.DrawCell(win, j*Size, (Rows-i-1)*Size, color.RGBA{A: 255})
			}
		}
	}

	win.Update()
}

func (b *Board) DrawCell(win *pixelgl.Window, x, y int, color color.RGBA) {
	rect := pixel.R(float64(x), float64(y), float64(x+Size), float64(y+Size))
	img := pixel.MakePictureData(pixel.R(0, 0, Size, Size))
	for i := range img.Pix {
		img.Pix[i] = color
	}

	sprite := pixel.NewSprite(img, img.Bounds())
	sprite.Draw(win, pixel.IM.Moved(rect.Center()).Scaled(rect.Center(), 0.5))
}

func (b *Board) Update() {
	newGrid := make(Grid, Rows)
	for i := range b.Grid {
		newGrid[i] = make([]mutation.Cell, Cols)
		copy(newGrid[i], b.Grid[i])

		for j := range b.Grid[i] {
			if isAlive(b.Grid[i][j]) {
				if b.Rules.UnderPopulation(b.Grid, i, j) || b.Rules.OverPopulation(b.Grid, i, j) {
					kill(&newGrid[i][j])
				}
				if isMutant(b.Grid[i][j]) && b.Rules.DieFromInstability(b.Grid, i, j) {
					kill(&newGrid[i][j])
				}
				if haveMutation(b.Grid[i][j], "Warrior cell") {
					for _, cell := range getAdjacentLivingCells(b.Grid, i, j, 2) {
						if b.Rules.WarriorInvasion() {
							kill(&newGrid[cell.I][cell.J])
						}
					}
				}
			} else {
				if canReproduce, _, mutantParents := b.Rules.Reproduce(b.Grid, i, j); canReproduce == true {
					create(&newGrid[i][j])
					if mutation.CanMutate(b.Grid[i][j], mutantParents) {
						mutation.FindMutation(&newGrid[i][j], mutantParents)
					}
				}
			}
		}
	}

	b.Grid = newGrid
}

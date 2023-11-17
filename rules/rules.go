package rules

import (
	"game-of-life/board"
)

type Rules struct{}

func (r *Rules) UnderPopulation(grid board.Grid, i, j int) bool {
	if grid[i][j].Mutation.Name == "Lonely Cell" {
		return false
	}
	return board.CountNeighbors(grid, i, j) < 2
}

func (r *Rules) OverPopulation(grid board.Grid, i, j int) bool {
	if grid[i][j].Mutation.Name == "Friendly Cell" {
		return board.CountNeighbors(grid, i, j) > 4
	}
	return board.CountNeighbors(grid, i, j) > 3
}

func (r *Rules) Reproduce(grid board.Grid, i, j int) bool {
	return board.CountNeighbors(grid, i, j) == 3
}

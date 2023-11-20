package rules

import (
	"game-of-life/board"
)

type Rules struct{}

func (r *Rules) UnderPopulation(grid board.Grid, i, j int) bool {
	if grid[i][j].Mutation.Name == "Lonely Cell" {
		return false
	}
	neighbors, _ := board.CountNeighbors(grid, i, j)
	return neighbors < 2
}

func (r *Rules) OverPopulation(grid board.Grid, i, j int) bool {
	if grid[i][j].Mutation.Name == "Friendly Cell" {
		neighbors, _ := board.CountNeighbors(grid, i, j)

		return neighbors > 4
	}
	neighbors, _ := board.CountNeighbors(grid, i, j)

	return neighbors > 3
}

func (r *Rules) Reproduce(grid board.Grid, i, j int) bool {
	neighbors, _ := board.CountNeighbors(grid, i, j)

	return neighbors == 3
}

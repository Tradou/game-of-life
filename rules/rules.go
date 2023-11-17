package rules

import "game-of-life/board"

type Rules struct{}

func (r *Rules) UnderPopulation(grid board.Grid, i, j int) bool {
	return board.CountNeighbors(grid, i, j) < 2
}

func (r *Rules) OverPopulation(grid board.Grid, i, j int) bool {
	return board.CountNeighbors(grid, i, j) > 3
}

func (r *Rules) Reproduce(grid board.Grid, i, j int) bool {
	return board.CountNeighbors(grid, i, j) == 3
}

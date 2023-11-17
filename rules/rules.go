package rules

import "game-of-life/board"

type Rules struct{}

func (r *Rules) Underpopulation(grid board.Grid, i, j int) bool {
	return grid[i][j] && board.CountNeighbors(grid, i, j) < 2
}

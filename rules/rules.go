package rules

import (
	"game-of-life/board"
	"math/rand"
)

type Rules struct{}

func (r *Rules) UnderPopulation(grid board.Grid, i, j int) bool {
	if grid[i][j].Mutation.Name == "Lonely Cell" {
		return false
	}
	if grid[i][j].Mutation.Name == "Warrior Cell" {
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

	if grid[i][j].Mutation.Name == "Warrior Cell" {
		return false
	}

	neighbors, _ := board.CountNeighbors(grid, i, j)

	return neighbors > 3
}

func (r *Rules) Reproduce(grid board.Grid, i, j int) (bool, int, int) {
	n, m := board.CountNeighbors(grid, i, j)

	return n == 3, n, m
}

func (r *Rules) DieFromInstability(grid board.Grid, i, j int) bool {
	return rand.Intn(100) >= grid[i][j].Mutation.Stability
}

func (r *Rules) WarriorInvasion() bool {
	return rand.Intn(100) <= 20
}

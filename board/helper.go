package board

import (
	"game-of-life/mutation"
)

func CountNeighbors(grid Grid, i, j int) (int, int) {
	neighbors := 0
	mNeighbors := 0

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adjacentI, adjacentJ := i+di, j+dj

			if IsInside(adjacentI, adjacentJ, len(grid), len(grid[0])) {
				if isAlive(grid[adjacentI][adjacentJ]) {
					if isMutant(grid[adjacentI][adjacentJ]) {
						mNeighbors++
					}
					neighbors++
				}
			}
		}
	}

	return neighbors, mNeighbors
}

func IsInside(i, j, rows, cols int) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols
}

func isAlive(c mutation.Cell) bool {
	return c.State
}

func isMutant(c mutation.Cell) bool {
	return isAlive(c) && c.Mutation.Name != ""
}

func haveMutation(c mutation.Cell, m string) bool {
	return isAlive(c) && c.Mutation.Name == m
}

func kill(c *mutation.Cell) {
	c.State = false
	c.Mutation = mutation.Attribute{}
}

func create(c *mutation.Cell) {
	c.State = true
}

func getAdjacentLivingCells(grid Grid, i, j, r int) []struct {
	I int
	J int
} {
	var adjacentCells []struct {
		I int
		J int
	}

	for di := -r; di <= r; di++ {
		for dj := -r; dj <= r; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adjacentI, adjacentJ := i+di, j+dj

			if IsInside(adjacentI, adjacentJ, len(grid), len(grid[0])) {
				if isAlive(grid[adjacentI][adjacentJ]) {
					adjacentCells = append(adjacentCells, struct {
						I int
						J int
					}{I: adjacentI, J: adjacentJ})
				}
			}
		}
	}

	return adjacentCells
}

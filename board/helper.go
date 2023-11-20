package board

import "game-of-life/mutation"

func CountNeighbors(grid Grid, i, j int) int {
	neighbors := 0

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adjacentI, adjacentJ := i+di, j+dj

			if IsInside(adjacentI, adjacentJ, len(grid), len(grid[0])) {
				if isAlive(grid[adjacentI][adjacentJ]) {
					neighbors++
				}
			}
		}
	}

	return neighbors
}

func IsInside(i, j, rows, cols int) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols
}

func isAlive(c mutation.Cell) bool {
	return c.State == "ALIVE"
}

func isMutant(c mutation.Cell) bool {
	return c.Mutation.Name != ""
}

func haveMutation(c mutation.Cell, m string) bool {
	return c.Mutation.Name == m
}

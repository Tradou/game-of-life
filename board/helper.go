package board

func CountNeighbors(grid Grid, i, j int) int {
	neighbors := 0

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adjacentI, adjacentJ := i+di, j+dj

			if IsInside(adjacentI, adjacentJ, len(grid), len(grid[0])) {
				if grid[adjacentI][adjacentJ] {
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

func isAlive(grid Grid, i, j int) bool {
	return grid[i][j]
}

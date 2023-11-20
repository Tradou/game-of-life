package board

import (
	"game-of-life/mutation"
	"testing"
)

func TestIsInside(t *testing.T) {
	tests := []struct {
		i, j     int
		rows     int
		cols     int
		isInside bool
	}{
		{0, 0, 4, 4, true},
		{1, 2, 3, 4, true},
		{0, 2, 3, 4, true},
		{2, 2, 3, 4, true},
		{1, 0, 3, 4, true},
		{1, 4, 3, 4, false},
		{-1, 2, 3, 4, false},
		{1, 5, 3, 4, false},
		{3, 2, 3, 4, false},
		{1, -1, 3, 4, false},
		{-1, 5, 3, 4, false},
		{3, 5, 3, 4, false},
		{3, -1, 3, 4, false},
		{-1, -1, 3, 4, false},
	}

	for _, tt := range tests {
		t.Run("IsInside fn", func(t *testing.T) {
			got := IsInside(tt.i, tt.j, tt.rows, tt.cols)
			if got != tt.isInside {
				t.Errorf("IsInside(%d, %d, %d, %d) = %v, want %v", tt.i, tt.j, tt.rows, tt.cols, got, tt.isInside)
			}
		})
	}
}

func TestCountNeighbors(t *testing.T) {
	tests := []struct {
		name          string
		grid          Grid
		row, col      int
		expectedAlive int
	}{
		{name: "NoNeighbors", grid: Grid{{{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "ALIVE"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}}, row: 1, col: 1},
		{name: "OneNeighbor", grid: Grid{{{State: "ALIVE"}, {State: "DEAD"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "ALIVE"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}}, row: 1, col: 1, expectedAlive: 1},
		{name: "FourNeighbors", grid: Grid{{{State: "ALIVE"}, {State: "DEAD"}, {State: "ALIVE"}}, {{State: "DEAD"}, {State: "ALIVE"}, {State: "DEAD"}}, {{State: "ALIVE"}, {State: "DEAD"}, {State: "ALIVE"}}}, row: 1, col: 1, expectedAlive: 4},
		{name: "AllNeighborsAlive", grid: Grid{{{State: "ALIVE"}, {State: "ALIVE"}, {State: "ALIVE"}}, {{State: "ALIVE"}, {State: "ALIVE"}, {State: "ALIVE"}}, {{State: "ALIVE"}, {State: "ALIVE"}, {State: "ALIVE"}}}, row: 1, col: 1, expectedAlive: 8},
		{name: "EdgeCase", grid: Grid{{{State: "ALIVE"}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountNeighbors(tt.grid, tt.row, tt.col)
			if got != tt.expectedAlive {
				t.Errorf("CountNeighbors(%v, %d, %d) = %d, want %d", tt.grid, tt.row, tt.col, got, tt.expectedAlive)
			}
		})
	}
}

func TestIsAlive(t *testing.T) {
	tests := []struct {
		name          string
		grid          Grid
		row, col      int
		expectedAlive bool
	}{
		{name: "BeDead", grid: Grid{{{State: "DEAD"}}}, row: 0, col: 0, expectedAlive: false},
		{name: "BeAlive", grid: Grid{{{State: "ALIVE"}}}, row: 0, col: 0, expectedAlive: true},
		{name: "UnknownCaseConsideredAsDead", grid: Grid{{{State: "FOO"}}}, row: 0, col: 0, expectedAlive: false},
		{name: "BeDead", grid: Grid{{{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "ALIVE"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}}, row: 1, col: 0, expectedAlive: false},
		{name: "BeAlive", grid: Grid{{{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "ALIVE"}, {State: "DEAD"}}, {{State: "DEAD"}, {State: "DEAD"}, {State: "DEAD"}}}, row: 1, col: 1, expectedAlive: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAlive(tt.grid[tt.row][tt.col])
			if got != tt.expectedAlive {
				t.Errorf("isAlive(%v, %d, %d) = %t, want %t", tt.grid, tt.row, tt.col, got, tt.expectedAlive)
			}
		})
	}
}

func TestIsMutant(t *testing.T) {
	tests := []struct {
		name           string
		grid           Grid
		row, col       int
		expectedMutant bool
	}{
		{name: "Cell is dead and not mutant", grid: Grid{{{State: "DEAD"}}}, row: 0, col: 0, expectedMutant: false},
		{name: "Cell is dead and mutant", grid: Grid{{{State: "DEAD", Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutant: true},
		{name: "Cell is alive and not mutant", grid: Grid{{{State: "ALIVE"}}}, row: 0, col: 0, expectedMutant: false},
		{name: "Cell is alive and mutant", grid: Grid{{{State: "ALIVE", Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutant: true},
		{name: "UnknownCaseConsideredAsMutant", grid: Grid{{{State: "ALIVE", Mutation: mutation.Attribute{Name: "TOTO MUTATION"}}}}, row: 0, col: 0, expectedMutant: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMutant(tt.grid[tt.row][tt.col])
			if got != tt.expectedMutant {
				t.Errorf("isMutant(%v, %d, %d) = %t, want %t", tt.grid, tt.row, tt.col, got, tt.expectedMutant)
			}
		})
	}
}

package board

import (
	"game-of-life/mutation"
	"github.com/stretchr/testify/assert"
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
		name                          string
		grid                          Grid
		row, col                      int
		expectedAlive, expectedMutant int
	}{
		{name: "NoNeighbors", grid: Grid{{{State: false}, {State: false}, {State: false}}, {{State: false}, {State: true}, {State: false}}, {{State: false}, {State: false}, {State: false}}}, row: 1, col: 1},
		{name: "OneNeighbor", grid: Grid{{{State: true}, {State: false}, {State: false}}, {{State: false}, {State: true}, {State: false}}, {{State: false}, {State: false}, {State: false}}}, row: 1, col: 1, expectedAlive: 1},
		{name: "FourNeighbors", grid: Grid{{{State: true}, {State: false}, {State: true}}, {{State: false}, {State: true}, {State: false}}, {{State: true}, {State: false}, {State: true}}}, row: 1, col: 1, expectedAlive: 4},
		{name: "AllNeighborsAlive", grid: Grid{{{State: true}, {State: true}, {State: true}}, {{State: true}, {State: true}, {State: true}}, {{State: true}, {State: true}, {State: true}}}, row: 1, col: 1, expectedAlive: 8},
		{name: "EdgeCase", grid: Grid{{{State: true}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNeighbors, gotMutants := CountNeighbors(tt.grid, tt.row, tt.col)
			if gotNeighbors != tt.expectedAlive {
				t.Errorf("CountNeighbors(%v, %d, %d) = %d, want %d", tt.grid, tt.row, tt.col, gotNeighbors, tt.expectedAlive)
			}
			if gotMutants != tt.expectedMutant {
				t.Errorf("CountNeighbors(%v, %d, %d) = %d, want %d", tt.grid, tt.row, tt.col, gotMutants, tt.expectedMutant)
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
		{name: "BeDead", grid: Grid{{{State: false}}}, row: 0, col: 0, expectedAlive: false},
		{name: "BeAlive", grid: Grid{{{State: true}}}, row: 0, col: 0, expectedAlive: true},
		{name: "BeDead", grid: Grid{{{State: false}, {State: false}, {State: false}}, {{State: false}, {State: true}, {State: false}}, {{State: false}, {State: false}, {State: false}}}, row: 1, col: 0, expectedAlive: false},
		{name: "BeAlive", grid: Grid{{{State: false}, {State: false}, {State: false}}, {{State: false}, {State: true}, {State: false}}, {{State: false}, {State: false}, {State: false}}}, row: 1, col: 1, expectedAlive: true},
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
		{name: "Cell is dead and not mutant", grid: Grid{{{State: false}}}, row: 0, col: 0, expectedMutant: false},
		{name: "Cell is dead and mutant", grid: Grid{{{State: false, Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutant: false},
		{name: "Cell is alive and not mutant", grid: Grid{{{State: true}}}, row: 0, col: 0, expectedMutant: false},
		{name: "Cell is alive and mutant", grid: Grid{{{State: true, Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutant: true},
		{name: "UnknownCaseConsideredAsMutant", grid: Grid{{{State: true, Mutation: mutation.Attribute{Name: "TOTO MUTATION"}}}}, row: 0, col: 0, expectedMutant: true},
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

func TestHaveMutation(t *testing.T) {
	mutationName := "Lonely Cell"

	tests := []struct {
		name             string
		grid             Grid
		row, col         int
		mutationName     string
		expectedMutation bool
	}{
		{name: "Cell is dead and not mutant", grid: Grid{{{State: false}}}, row: 0, col: 0, expectedMutation: false},
		{name: "Cell is dead and mutant", grid: Grid{{{State: false, Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutation: false},
		{name: "Cell is alive and not mutant", grid: Grid{{{State: true}}}, row: 0, col: 0, expectedMutation: false},
		{name: "Cell is alive and mutant", grid: Grid{{{State: true, Mutation: mutation.Attribute{Name: "Lonely Cell"}}}}, row: 0, col: 0, expectedMutation: true},
		{name: "UnknownCaseConsideredAsMutant", grid: Grid{{{State: true, Mutation: mutation.Attribute{Name: "TOTO MUTATION"}}}}, row: 0, col: 0, expectedMutation: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := haveMutation(tt.grid[tt.row][tt.col], mutationName)
			if got != tt.expectedMutation {
				t.Errorf("haveMutation(%v, %d, %d) = %t, want %t", tt.grid, tt.row, tt.col, got, tt.expectedMutation)
			}
		})
	}
}

func TestGetAdjacentCells(t *testing.T) {
	type cells []struct {
		I int
		J int
	}

	tests := []struct {
		name          string
		grid          Grid
		row, col      int
		radius        int
		expectedCells cells
	}{
		{name: "Empty grid", grid: Grid{{}}, row: 0, col: 0, radius: 1, expectedCells: cells{}},
		{name: "Grid with one cell not alive", grid: Grid{{{State: false}}}, row: 0, col: 0, radius: 1, expectedCells: cells{}},
		{name: "Grid with only one cell alive", grid: Grid{{{State: true}}}, row: 0, col: 0, radius: 1, expectedCells: cells{}},
		{name: "Grid with one cell alive", grid: Grid{{{State: true}, {State: false}}}, row: 0, col: 0, radius: 1, expectedCells: cells{}},
		{name: "Grid with three cells alive", grid: Grid{{{State: true}, {State: false}}, {{State: true}, {State: true}}}, row: 0, col: 1, radius: 1, expectedCells: cells{{I: 0, J: 0}, {I: 1, J: 0}, {I: 1, J: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAdjacentLivingCells(tt.grid, tt.row, tt.col, tt.radius)
			if !assert.ElementsMatch(t, got, tt.expectedCells) {
				t.Errorf("getAdjacentLivingCells(%v, %d, %d, %d) = %v, want %v", tt.grid, tt.row, tt.col, tt.radius, got, tt.expectedCells)
			}
		})
	}
}

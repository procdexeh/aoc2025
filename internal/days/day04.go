package days

import (
	"strings"
)

type Day04 struct{}

func (d *Day04) Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := [][]string{}

	positions := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	sum := 0
	for _, v := range lines {
		row := strings.Split(v, "")
		grid = append(grid, row)
	}

	for col_i, col := range grid {
		for i, v := range col {
			if v == "@" {
				adj_rolls := 0

				for _, position := range positions {
					x := position[0]
					y := position[1]

					check_x := x + i
					check_y := y + col_i

					if check_x < 0 || check_x >= len(col) {
						continue
					}

					if check_y < 0 || check_y >= len(grid) {
						continue
					}

					if grid[check_y][check_x] == "@" {
						adj_rolls += 1
					}

				}
				if adj_rolls < 4 {
					sum += 1
				}
			}
		}
	}

	return sum
}

func (d *Day04) Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := [][]string{}

	positions := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	sum := 0
	for _, v := range lines {
		row := strings.Split(v, "")
		grid = append(grid, row)
	}

	// this is really bruteforce; just check if the sum changes for the while loop, write roll deletions to copy of grid, and swap them out if we're going again
	for {
		changed := sum
		pending_grid := grid
		for col_i, col := range grid {
			for i, v := range col {
				if v == "@" {
					adj_rolls := 0

					for _, position := range positions {
						x := position[0]
						y := position[1]

						check_x := x + i
						check_y := y + col_i

						if check_x < 0 || check_x >= len(col) {
							continue
						}

						if check_y < 0 || check_y >= len(grid) {
							continue
						}

						if grid[check_y][check_x] == "@" {
							adj_rolls += 1
						}

					}
					if adj_rolls < 4 {
						sum += 1
						pending_grid[col_i][i] = "x"
					}
				}
			}
		}

		if changed == sum {
			break
		}

		grid = pending_grid
	}

	return sum
}

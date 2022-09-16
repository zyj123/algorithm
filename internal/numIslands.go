package internal

func numIslands(grid [][]byte) int {
	r, c := len(grid), len(grid[0])
	ret := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] != '1' {
				continue
			}
			ret++
			dyeing(grid, i, j)
		}
	}
	return ret
}

func dyeing(grid [][]byte, i, j int) {
	r, c := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= r || j >= c {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	borders := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, border := range borders {
		dyeing(grid, i+border[0], j+border[1])
	}
}

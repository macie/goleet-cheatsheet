package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == '1' && !visited[i][j] {
				dfsIslands(grid, visited, i, j)
				count++
			}
		}
	}
	return count
}

func dfsIslands(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfsIslands(grid, visited, i-1, j)
	dfsIslands(grid, visited, i+1, j)
	dfsIslands(grid, visited, i, j-1)
	dfsIslands(grid, visited, i, j+1)
}

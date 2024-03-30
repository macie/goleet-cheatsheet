package main

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)
	if len(heights) == 0 || len(heights[0]) == 0 {
		return res
	}
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
	}
	atlantic := make([][]bool, m)
	for i := range atlantic {
		atlantic[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dfsOceans(heights, pacific, -1, i, 0)
		dfsOceans(heights, atlantic, -1, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfsOceans(heights, pacific, -1, 0, i)
		dfsOceans(heights, atlantic, -1, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfsOceans(heights [][]int, visit [][]bool, pre, i, j int) {
	m, n := len(heights), len(heights[0])
	if i < 0 || i >= m || j < 0 || j >= n ||
		visit[i][j] || pre > heights[i][j] {
		return
	}
	visit[i][j] = true
	dfsOceans(heights, visit, heights[i][j], i, j+1)
	dfsOceans(heights, visit, heights[i][j], i, j-1)
	dfsOceans(heights, visit, heights[i][j], i+1, j)
	dfsOceans(heights, visit, heights[i][j], i-1, j)
}

func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
	n := len(grid[0])

	max_area_of_island := 0

	for i := 0; i<m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == 1{
				max_area_of_island = max(max_area_of_island, dfs(grid, i, j))
			}
		}
	}
	return max_area_of_island
}

func dfs(grid [][]int, i, j int)int{
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0{
		return 0
	}else{
		grid[i][j] = 0
		return 1 + dfs(grid, i, j+1) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i-1, j)
	} 
}

func max(a, b int)int{
	if a > b {
		return a
	}
	return b
}


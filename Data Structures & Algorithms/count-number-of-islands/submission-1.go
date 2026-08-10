func numIslands(grid [][]byte) int {
    m := len(grid)
	n := len(grid[0])
	num_of_country := 0

	for i := 0; i<m; i++{
		for j:= 0; j<n; j++{
			if grid[i][j] == '1'{
				num_of_country ++
				dfs(grid, m, n, i, j)
			}
		}
	} 
	return num_of_country
}


func dfs(grid [][]byte, m, n, i, j int){
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0'{
		return
	}else{
		grid[i][j] = '0'
		dfs(grid, m, n, i, j+1)
		dfs(grid, m, n, i, j-1)
		dfs(grid, m, n, i+1, j)
		dfs(grid, m, n, i-1, j)
	}
}
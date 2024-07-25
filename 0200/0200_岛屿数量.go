package numberOfIslands

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	res := 0
	nr := len(grid)
	nc := len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		grid[r][c] = '0'
		if r-1 >= 0 && grid[r-1][c] == '1' {
			dfs(r-1, c)
		}
		if r+1 < nr && grid[r+1][c] == '1' {
			dfs(r+1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == '1' {
			dfs(r, c-1)
		}
		if c+1 < nc && grid[r][c+1] == '1' {
			dfs(r, c+1)
		}
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				dfs(r, c)
				res++
			}
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)

package rottingOranges

//leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {
	had := false
	hadFresh := false

	nr := len(grid)
	nc := len(grid[0])

	i := 2
	for ; ; i++ {
		for r := 0; r < nr; r++ {
			for c := 0; c < nc; c++ {
				if grid[r][c] == i {
					if r-1 >= 0 && grid[r-1][c] == 1 {
						grid[r-1][c] = i + 1
						had = true
					}
					if r+1 < nr && grid[r+1][c] == 1 {
						grid[r+1][c] = i + 1
						had = true
					}
					if c-1 >= 0 && grid[r][c-1] == 1 {
						grid[r][c-1] = i + 1
						had = true
					}
					if c+1 < nc && grid[r][c+1] == 1 {
						grid[r][c+1] = i + 1
						had = true
					}
				}

				if grid[r][c] == 1 {
					hadFresh = true
				}
			}
		}

		if had {
			had = false
			hadFresh = false
		} else {
			break
		}
	}

	if hadFresh {
		return -1
	}

	return i - 2
}

//leetcode submit region end(Prohibit modification and deletion)

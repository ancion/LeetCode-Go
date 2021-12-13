package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var topBottomSkyline []int
	var leftRightSKyline []int
	for i := range grid {
		cur := 0
		for _, v := range grid[i] {
			if cur < v {
				cur = v
			}
		}
		leftRightSKyline = append(leftRightSKyline, cur)
	}
	for j := range grid {
		cur := 0
		for i := 0; i < len(grid[0]); i++ {
			if cur < grid[i][j] {
				cur = grid[i][j]
			}
		}
		topBottomSkyline = append(topBottomSkyline, cur)
	}
	var ans int
	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			ans += min(topBottomSkyline[j], leftRightSKyline[i]) - grid[i][j]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

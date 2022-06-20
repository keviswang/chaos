package matrix

// 设定边界
func generateMatrix(n int) [][]int {
	num := 1
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	up, down, left, right := 0, n-1, 0, n-1
	for num <= n*n {
		for i := left; i <= right; i++ {
			ret[up][i] = num
			num++
		}
		up++
		for i := up; i <= down; i++ {
			ret[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			ret[down][i] = num
			num++
		}
		down--
		for i := down; i >= up; i-- {
			ret[i][left] = num
			num++
		}
		left++
	}
	return ret
}

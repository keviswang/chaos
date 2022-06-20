package matrix

func TraverseOrder1(input [][]int) []int {
	var result []int
	if len(input) == 0 {
		return result
	}
	currentRow, maxRow, currentCol, maxCol := 0, len(input)-1, 0, len(input[0])-1
	for {
		// 向右
		for i := currentCol; i <= maxCol; i++ {
			result = append(result, input[currentRow][i])
		}
		currentRow++
		if currentRow > maxRow {
			break
		}
		// 向下
		for i := currentRow; i <= maxRow; i++ {
			result = append(result, input[i][maxCol])
		}
		maxCol--
		if maxCol < currentCol {
			break
		}
		// 向左
		for i := maxCol; i >= currentCol; i-- {
			result = append(result, input[maxRow][i])
		}
		maxRow--
		if maxRow < currentRow {
			break
		}
		// 向上
		for i := maxRow; i >= currentRow; i-- {
			result = append(result, input[i][currentCol])
		}
		currentCol++
		if currentCol > maxCol {
			break
		}
	}
	return result
}

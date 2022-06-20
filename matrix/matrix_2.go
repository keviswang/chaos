package matrix

const (
	// RightDirection 方向向右
	RightDirection = 1
	// DownDirection 方向向下
	DownDirection = 2
	// LeftDirection 方向向左
	LeftDirection = 3
	// UpDirection 方向向上
	UpDirection = 4
)

type Iter struct {
	input [][]int
	// 索引
	index int
	// currentRow 当前行
	currentRow int
	// currentCol 当前列
	currentCol int
	// 最大行
	maxRow int
	// 最大列
	maxCol int
	// 方向
	direction int
}

func NewIter(input [][]int) *Iter {
	return &Iter{
		input:      input,
		index:      0,
		currentRow: 0,
		currentCol: 0,
		maxRow:     len(input) - 1,
		maxCol:     len(input[0]) - 1,
		direction:  RightDirection,
	}
}

func (i *Iter) Next() (bool, int) {
	if i.direction == RightDirection {
		c := i.index
		if c <= i.maxCol {
			i.index++
			return true, i.input[i.currentRow][c]
		}
		i.currentRow++
		if i.currentRow > i.maxRow {
			return false, 0
		}
		i.direction = DownDirection
		i.index = i.currentRow
		i.index++
		return true, i.input[i.currentRow][i.maxCol]
	} else if i.direction == DownDirection {
		r := i.index
		if r <= i.maxRow {
			i.index++
			return true, i.input[r][i.maxCol]
		}
		i.maxCol--
		if i.maxCol < i.currentCol {
			return false, 0
		}
		i.direction = LeftDirection
		i.index = i.maxCol
		i.index--
		return true, i.input[i.maxRow][i.maxCol]
	} else if i.direction == LeftDirection {
		c := i.index
		if c >= i.currentCol {
			i.index--
			return true, i.input[i.maxRow][c]
		}
		i.maxRow--
		if i.maxRow < i.currentRow {
			return false, 0
		}
		i.direction = UpDirection
		i.index = i.maxRow
		i.index--
		return true, i.input[i.maxRow][i.currentCol]
	} else if i.direction == UpDirection {
		r := i.index
		if r >= i.currentRow {
			i.index--
			return true, i.input[r][i.currentCol]
		}
		i.currentCol++
		if i.currentCol > i.maxCol {
			return false, 0
		}
		i.direction = RightDirection
		i.index = i.currentCol
		i.index++
		return true, i.input[i.currentRow][i.currentCol]
	}

	return false, 0
}

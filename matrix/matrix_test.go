package matrix

import (
	"fmt"
	"testing"
)

func TestTraverseOrder(t *testing.T) {

	matrix := generateMatrix(10)
	fmt.Println(TraverseOrder1(matrix))

	iter := NewIter(matrix)
	var result []int
	for {
		ok, value := iter.Next()
		if ok {
			result = append(result, value)

		} else {
			break
		}
	}
	fmt.Println(result)
}

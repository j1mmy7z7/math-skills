package skills

import (
	"math/big"
)

// quicksort algorithm for sorting the array
func QuickSort(array []*big.Float, start, end int) []*big.Float {
	if start >= end {
		return array
	}
	pivotIndex := partition(array, start, end)
	QuickSort(array, start, pivotIndex-1)
	QuickSort(array, pivotIndex+1, end)
	return array
}

func partition(array []*big.Float, start int, end int) int {
	pivot := array[end]
	i := start - 1

	for j := start; j < end; j++ {
		if array[j].Cmp(pivot) == -1 {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	i++
	array[i], array[end] = array[end], array[i]
	return i
}

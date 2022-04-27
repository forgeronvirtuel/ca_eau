package water

// BubbleSort implement the bubble sort algorithm.
func BubbleSort(toSort []int) []int {
	for cptr := len(toSort) - 1; cptr > 0; cptr-- {
		for i := 0; i < cptr; i++ {
			if toSort[i] > toSort[i+1] {
				toSort[i], toSort[i+1] = toSort[i+1], toSort[i]
			}
		}
	}
	return toSort
}

// SelectionSort implement the selection sort algorithm.
func SelectionSort(toSort []int) []int {
	for i1 := range toSort {
		min := i1
		for i2 := i1; i2 < len(toSort); i2++ {
			if toSort[i2] < toSort[min] {
				min = i2
			}
		}
		toSort[i1], toSort[min] = toSort[min], toSort[i1]
	}
	return toSort
}

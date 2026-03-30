package sort

func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, right, equal := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}
	QuickSort(left)
	QuickSort(right)

	copy(arr, append(append(left, equal...), right...))

	return arr
}

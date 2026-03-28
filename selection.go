package sort

/*

Selection Sort

Complexity: O(n²)
Space complexity: O(1)

The idea is to sequentially search for the minimum element
in the unsorted portion and move it to the beginning.
The algorithm is complex and is only effective for small lists.
*/
func SelectionSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n-1; i++ {

		minIdx := i
		for j := i + 1; j < n; j++ {

			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr
}

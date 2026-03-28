package sort

/*
	Bubble Sort

	Complexity: O(n²)
	Space Complexity: O(1)

	A sequential comparison and exchange of adjacent elements,
	in which the largest values ​​("heavy") "sink" to the end of the array,
	and the smaller ("light") ones float to the beginning
*/
func BubbleSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n-1; i++ {

		swapped := false

		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

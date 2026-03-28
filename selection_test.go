package sort

import (
	"reflect"
	"testing"
)

func Test_SelectionSort(t *testing.T) {

	tests := []struct {
		name        string
		description string
		data        []int
		want        []int
	}{
		{
			name:        "1",
			description: "nil",
			data:        []int{},
			want:        []int{},
		},
		{
			name:        "2",
			description: "default",
			data:        []int{2, 3, 0},
			want:        []int{0, 2, 3},
		},
		{
			name:        "3",
			description: "double",
			data:        []int{1, 1, 2},
			want:        []int{1, 1, 2},
		},
		{
			name:        "4",
			description: "negative values",
			data:        []int{1, -1, 2},
			want:        []int{-1, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			data := SelectionSort(test.data)
			if !reflect.DeepEqual(data, test.want) {
				t.Fatalf("have - %v, need - %v", data, test.want)
			}
		})
	}

}

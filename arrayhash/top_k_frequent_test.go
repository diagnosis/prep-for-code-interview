package arrayhash

import (
	"slices"
	"testing"
)

func Test_TopKFrequent(t *testing.T) {
	impls := []struct {
		name   string
		method func([]int, int) []int
	}{
		{name: "bucket sort", method: topKFrequent},
	}
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{name: "classic", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expected: []int{1, 2}},
		{"gap in frequencies", []int{1, 1, 1, 1, 1, 2}, 2, []int{1, 2}},
		{"shared frequency", []int{1, 1, 1, 2, 2, 3, 3}, 3, []int{1, 2, 3}},
		{"zero is a value", []int{0, 0, 0, 1}, 1, []int{0}},
		{"k equals distinct count", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"negatives", []int{-1, -1, 2}, 1, []int{-1}},
	}

	for _, impl := range impls {
		for _, tt := range tests {
			name := tt.name + " with " + impl.name
			t.Run(name, func(t *testing.T) {
				actual := impl.method(tt.nums, tt.k)
				slices.Sort(actual)
				slices.Sort(tt.expected)
				if !slices.Equal(tt.expected, actual) {
					t.Errorf("expected %v got %v", tt.expected, actual)
				}
			})
		}
	}

}

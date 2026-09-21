package arrayhash

import (
	"slices"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	impls := []struct {
		name   string
		method func([]int, int) []int
	}{
		{
			name:   "map",
			method: twoSum,
		},
	}
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "TC:1- happy", nums: []int{3, 4, 5, 6}, target: 7, expected: []int{0, 1}},
		{name: "TC:2- happy 2", nums: []int{4, 5, 6}, target: 10, expected: []int{0, 2}},
		{name: "TC:3-happy 2 same val", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{name: "TC:4-unhappy", nums: []int{1, 23, 4}, target: 3, expected: nil},
		{name: "target is double an element", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "negatives", nums: []int{-3, 4, 3, 90}, target: 0, expected: []int{0, 2}},
		{name: "answer at far ends", nums: []int{2, 7, 11, 15, 1, 8}, target: 10, expected: []int{0, 5}},
		{name: "zeros", nums: []int{0, 4, 3, 0}, target: 0, expected: []int{0, 3}},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			name := tt.name + " with " + impl.name
			t.Run(name, func(t *testing.T) {
				actual := impl.method(tt.nums, tt.target)
				if !slices.Equal(tt.expected, actual) {
					t.Errorf("expected two sum returns %v but got %v", tt.expected, actual)
				}
			})
		}
	}
}

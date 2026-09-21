package arrayhash

import (
	"fmt"
	"slices"
	"testing"
)

func Test_ContainsDuplicate(t *testing.T) {
	impls := []struct {
		name   string
		method func([]int) bool
	}{
		{"sort", containsDuplicateWithSort},
		{"map", containsDuplicate},
	}
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"TC-1:duplicate at end", []int{1, 3, 4, 5, 6, 3}, true},
		{"TC-2:all unique", []int{1, 3, 4, 5, 6}, false},
		{"TC-3:empty slice", []int{}, false},
		{"TC-4:negative duplicate", []int{-11, 2, 3, 4, 5, -11}, true},
		{"TC-5:single element", []int{5}, false},
		{"TC-6:all same", []int{5, 5, 5, 5, 5, 5, 5}, true},
		{"TC-7:adjacent duplicate", []int{1000, 10001, 1003, 1003, 1004}, true},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s with %s", tt.name, impl.name), func(t *testing.T) {
				in := slices.Clone(tt.input)
				actual := impl.method(in)
				if actual != tt.expected {
					t.Errorf("expected %v contains duplicate %v but got %v", tt.input, tt.expected, actual)
				}
			})
		}
	}
}

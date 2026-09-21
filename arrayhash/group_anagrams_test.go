package arrayhash

import (
	"slices"
	"testing"
)

func Test_GroupAnagrams(t *testing.T) {
	impls := []struct {
		name   string
		method func([]string) [][]string
	}{
		{"map", groupAnagrams},
	}
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{"classic", []string{"eat","tea","tan","ate","nat","bat"}, [][]string{{"eat","tea","ate"},{"tan","nat"},{"bat"}}},
		}
	for _, impl := range impls {
		for _, tt := range tests {
			name := tt.name + " with " + impl.name
			t.Run(name, func(t *testing.T) {
				actual := impl.method(tt.strs)
				normalize(actual)
				normalize(tt.expected)
				if !slices.EqualFunc(tt.expected, actual, slices.Equal) {
					t.Errorf("expected %v got %v", tt.expected, actual)
				}
			})
		}
	}
}

func normalize(strs [][]string) {
	for _, s := range strs {
		slices.Sort(s)
	}
	slices.SortFunc(strs, func(a, b []string) int {
		return slices.Compare(a, b)
	})
}

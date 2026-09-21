package arrayhash

import "testing"

func Test_ValidAnagram(t *testing.T) {
	impls := []struct {
		name   string
		method func(string, string) bool
	}{
		{name: "map", method: isAnagram},
		{name: "26Array", method: isAnagram26Int},
	}
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"TC:1-valid", "racecar", "carrace", true},
		{"TC:2-different length", "mahmut", "tuncera", false},
		{"TC:3-same chars different count", "saalihhabi", "sallihhabi", false},
		{"TC:4-empty", "", "", true},
		{"TC:5-identical", "safa", "safa", true},
		{"TC:2-different length swap", "tuncera", "mahmut", false},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			name := tt.name + " with " + impl.name
			t.Run(name, func(t *testing.T) {
				actual := impl.method(tt.s, tt.t)
				if tt.expected != actual {
					t.Errorf("expected %s and %s anagram to be %v got %v", tt.s, tt.t, tt.expected, actual)
				}
			})
		}
	}
}

package arrayhash

import (
	"slices"
	"testing"
)

func Test_Encode(t *testing.T) {
	s := &Solution{}
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "classic", input: []string{"Hello", "World"}, expected: "5#Hello5#World"},
		{name: "test", input: []string{"a#b", "cd"}, expected: "3#a#b2#cd"},
		{name: "empty string in list", input: []string{"", "a"}, expected: "0#1#a"},
		{name: "single empty string", input: []string{""}, expected: "0#"},
		{name: "empty list", input: []string{}, expected: ""},
		{name: "digits and hash", input: []string{"12#34", "5"}, expected: "5#12#341#5"},
		{name: "unicode", input: []string{"ş", "日本", "🙂"}, expected: "2#ş6#日本4#🙂"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := s.Encode(tt.input)
			if tt.expected != encoded {
				t.Fatalf("expected %s got %s", tt.expected, encoded)
			}
			decoded := s.Decode(encoded)
			if !slices.Equal(tt.input, decoded) {
				t.Fatalf("expected %v got %v", tt.input, decoded)
			}
		})
	}
}

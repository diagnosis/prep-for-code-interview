package arrayhash

import (
	"slices"
	"testing"
)

func Test_ProductOfArrayExceptSelf(t *testing.T){
	tests := []struct{
		name string
		input []int
		expected []int
	}{
		{name:"classic", input: []int{1,2,4,6}, expected: []int{48,24,12,8}},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			actual := productExceptSelf(tt.input)
			if !slices.Equal(tt.expected, actual){
				t.Errorf("expected %v got %v", tt.expected, actual)
			}
		})
	}

}
package arrayhash

import "slices"

func containsDuplicate(in []int) bool {
	m := make(map[int]struct{})
	for _, n := range in {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}

	return false
}
func containsDuplicateWithSort(in []int)bool{
	slices.Sort(in)
	for i := 1 ; i < len(in); i++{
		if in[i-1] == in[i]{
			return true
		}
	}
	return false
}

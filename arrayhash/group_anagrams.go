package arrayhash

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	res := make([][]string, 0)
	for _, str := range strs {
		array26 := [26]int{}
		for i := 0; i < len(str); i++ {
			array26[str[i]-'a']++
		}
		m[array26] = append(m[array26], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

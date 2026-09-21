package arrayhash

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if _, ok := ms[t[i]]; !ok {
			return false
		}
		ms[t[i]]--
		if ms[t[i]] == 0 {
			delete(ms, t[i])
		}
	}
	return true
}

func isAnagram26Int(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	array := [26]int{}
	for i := 0; i < len(s); i++ {
		array[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		array[t[i]-'a']--
	}
	if array == [26]int{} {
		return true
	}
	return false
}

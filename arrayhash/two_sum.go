package arrayhash

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if v, ok := m[c]; ok {
			return []int{v, i}

		}
		m[nums[i]] = i
	}
	return nil
}

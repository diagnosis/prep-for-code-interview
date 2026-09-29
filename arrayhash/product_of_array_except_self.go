package arrayhash

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	leftProd := make([]int, len(nums))
	rightProd := make([]int, len(nums))

	leftProd[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProd[i] = nums[i-1] * leftProd[i-1]
	}

	rightProd[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProd[i] = nums[i+1] * rightProd[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = leftProd[i] * rightProd[i]
	}
	return res
}

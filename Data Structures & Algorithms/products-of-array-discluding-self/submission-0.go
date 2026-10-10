func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	prefix := 1
	leftOutput := make([]int, n)
	for i := 0; i < n; i++ {
		leftOutput[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	rightOutput := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		rightOutput[i] = postfix
		postfix *= nums[i]
	} 

	for i := 0; i < n; i++ {
		output[i] = leftOutput[i] * rightOutput[i]
	}

	return output
}

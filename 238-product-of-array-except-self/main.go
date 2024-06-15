func productExceptSelf(nums []int) []int {
    
    n := len(nums)

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

    prefix := 1
	for i := 0; i < n; i++ {
		ans[i] *= prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= nums[i]
	}

	return ans
}

package main

func removeDuplicates(nums []int) int {
 	
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	var i int = 0
	for j := i; j < n; j++ {
	   if nums[j] != nums[i]{
		i += 1
		nums[i] = nums[j]
	   }

	}

	return i + 1


}

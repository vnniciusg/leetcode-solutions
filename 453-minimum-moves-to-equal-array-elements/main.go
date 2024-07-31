package main

func minMoves(nums []int) int {
    return sum(nums) + (len(nums) * min(nums))
}

func sum(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func min(nums []int) int {
	var min int = nums[0]
	for _, num := range nums[1:] {
		if min > num {
			min = num
		}
	}
	return min
}
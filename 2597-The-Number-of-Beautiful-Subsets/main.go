package main

import "sort"

func beautifulSubsets(nums []int, k int) int {

	count := 0

	var backtrack func(start int, subset []int)
	backtrack = func(start int, subset []int) {

		if subset != nil {
			count += 1
		}

		for i := start; i < len(nums); i++ {
			var canAdd bool = true

			for _, num := range subset {
				if abs(num-nums[i]) == k {
					canAdd = false
					break
				}
			}

			if canAdd {
				subset = append(subset, nums[i])
				backtrack(i+1, subset)
				subset = subset[:len(subset)-1]
			}
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrack(0, []int{})
	return count - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

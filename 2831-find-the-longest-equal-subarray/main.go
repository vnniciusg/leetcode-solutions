package main

import "fmt"

func main() {
	testArr1 := []int{1, 3, 2, 3, 1, 3}
	kArr1 := 3

	testArr2 := []int{1, 1, 2, 2, 1, 1}
	kArr2 := 2

	fmt.Printf("Solution 1: %d\n", longestEqualSubarray(testArr1, kArr1))
	fmt.Printf("Solution 2: %d\n", longestEqualSubarray(testArr2, kArr2))
}

func longestEqualSubarray(nums []int, k int) int {

	n := len(nums)
	j, res, cnt := 0, 0, 0

	freq := make(map[int]int)
	for i := 0; i < n; i++ {
		freq[nums[i]] += 1

		cnt = max(res, freq[nums[i]])

		if cnt+k < i-j+1 {
			freq[nums[j]]--
			j++
		} else {
			res = max(res, cnt)
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

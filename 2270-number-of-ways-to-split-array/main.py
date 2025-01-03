class Solution:
    def waysToSplitArray(self, nums: List[int]) -> int:
        
        left = 0
        right = sum(nums)

        count = 0
        n = len(nums)
        for idx in range(n - 1):
            left += nums[idx]
            right -= nums[idx]

            if left >= right: count += 1

        return count


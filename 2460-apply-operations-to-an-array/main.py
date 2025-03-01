class Solution:
    def applyOperations(self, nums: List[int]) -> List[int]:
        
        N = len(nums)

        for i in range(N - 1):
            if nums[i] == nums[i + 1]:
                nums[i] *= 2
                nums[i + 1] = 0

        l = 0
        for r in range(N):
            if nums[r]:
                nums[r], nums[l] = nums[l], nums[r]
                l += 1

        return nums

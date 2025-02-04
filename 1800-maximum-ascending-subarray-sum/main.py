class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        _max = 0
        curr_sub_sum = nums[0]

        for curr_idx in range(1, len(nums)):

            if nums[curr_idx] <= nums[curr_idx - 1]:
                _max = max(_max, curr_sub_sum)
                curr_sub_sum = 0 

            curr_sub_sum += nums[curr_idx]

        return max(_max, curr_sub_sum)

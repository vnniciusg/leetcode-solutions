class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        
        for row in grid:
            for n in row:
                if n % x != grid[0][0] % x:
                    return -1
        
        nums = sorted([n for row in grid for n in row])
        prefix, total = 0, sum(nums)
        res = float('inf')

        for i in range(len(nums)):
            cost_left = nums[i] * i - prefix
            cost_right = total - prefix - (nums[i] * (len(nums) - i))
            operations = (cost_left + cost_right) // x
            res = min(res, operations)
            prefix += nums[i]

        return res

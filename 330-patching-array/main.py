class Solution:
    def minPatches(self, nums: List[int], n: int) -> int:
        
        _sum, cnt, i = 0, 0, 0

        while _sum < n:
            if i < len(nums) and nums[i] <= _sum + 1:
                _sum += nums[i]
                i += 1
            else:
                cnt += 1
                _sum += _sum + 1
        
        return cnt

class Solution:
    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        
        nums.sort()

        def count_pairs(mid: int) -> int:
            n = len(nums)
            count, left = 0, 0

            for right in range(n):
                while nums[right] - nums[left] > mid:
                    left += 1
                count += right -left
            
            return count
        
        left, right = 0, max(nums) - min(nums)
        while left < right:
            mid = left + (right-left) // 2
            if count_pairs(mid) < k:
                left = mid + 1
            else:
                right = mid
        
        return left

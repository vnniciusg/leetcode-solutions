class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        
        n = len(nums)

        def aux(k: int) -> bool:

            _track = [0] * (n + 1)

            for i in range(k):
                left, right, val = queries[i]

                _track[left] += val
                _track[right + 1] -= val
            
            for i in range(1, n + 1):
                _track[i] += _track[i - 1]

            return all(_track[i] >= nums[i] for i in range(n))

        if sum(nums) == 0: return 0

        left, right = 1, len(queries)
        if not aux(right):
            return -1

        while left < right:
            mid = left + (right - left) // 2

            if aux(mid):
                right = mid
            else:
                left = mid + 1
        
        return left

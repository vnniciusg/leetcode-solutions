class Solution:
    def maximumCandies(self, candies: List[int], k: int) -> int:
        
        left, right = 1, max(candies)
        ans = 0

        while left <= right:
            mid = left + (right - left) // 2

            _count = sum(pile // mid for pile in candies)
            if _count >= k:
                ans = mid
                left = mid + 1
            else:
                right = mid - 1
            
        return ans

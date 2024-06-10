class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        
        n = len(heights)

        expected = sorted(heights.copy())
        
        return sum([1 for i in range(n) if heights[i] != expected[i] ])

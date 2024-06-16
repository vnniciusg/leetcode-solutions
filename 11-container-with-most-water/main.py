class Solution:
    def maxArea(self, height: List[int]) -> int:
        
        max_amount = 0
        i, j = 0, len(height) - 1

        while i < j:
            curr_amount = min(height[i], height[j]) * (j - i)

            if max_amount < curr_amount:
                max_amount = curr_amount
            elif height[i] < height[j]:
                i += 1
            else:
                j -= 1
        

        return max_amount

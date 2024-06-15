class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:           
        return [kid + extraCandies >= max(candies) for kid in candies]

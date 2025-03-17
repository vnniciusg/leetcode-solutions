class Solution:
    def divideArray(self, nums: List[int]) -> bool:
        return all(nums.count(num) & 1 == 0 for num in set(nums))

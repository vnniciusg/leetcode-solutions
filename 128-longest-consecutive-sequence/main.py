class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        
        nums_set = set(nums)
        max_length: int = 0

        for n in nums_set:
            if (n - 1) not in nums_set:
                length = 1
                while (n + length) in nums_set:
                    length += 1
                max_length = max(length, max_length)
        
        return max_length

from typing import List

class Solution:
    def beautifulSubsets(self, nums: List[int], k: int) -> int:
        def backtrack(start: int ,subset: List[int]) -> None:
            nonlocal count

            if subset:
                count += 1
                        
            for i in range(start, len(nums)):
                if all(abs(num - nums[i]) != k for num in subset):
                    backtrack(i + 1, subset + [nums[i]])
        
        nums.sort()
        count = 0
        backtrack(0, [])
        return count
        
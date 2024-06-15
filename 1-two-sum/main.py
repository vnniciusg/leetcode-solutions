class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        
        _hashmap = {}

        for idx, num in enumerate(nums):
            if _hashmap.get(num) is not None:
                return [idx, _hashmap.get(num)]
            s[target-num] = idx

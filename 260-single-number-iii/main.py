class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:

        d = defaultdict(int)
        for num in nums:
            d[num] = d.get(num, 0) + 1

        res = [k for k, v in d.items() if v == 1]
    
        return res

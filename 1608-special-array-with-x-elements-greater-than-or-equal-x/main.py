class Solution:
    def specialArray(self, nums: List[int]) -> int:
        
        nums.sort()
        n = len(nums)

        for x in range(1, n+1):
            count = 0
            for num in nums:
                if num >= x:
                    count += 1
            if count == x:
                return x
        
        return -1

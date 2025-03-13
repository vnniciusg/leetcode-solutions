class Solution:
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        
        n = len(nums)
        ans = [0] * (n + 1)

        for l, r in queries:
            ans[l] += 1
            ans[r + 1] -= 1

        for i in range(1, n  + 1):
            ans[i] += ans[i - 1]
                
        return all(ans[i] >= nums[i] for i in range(n))

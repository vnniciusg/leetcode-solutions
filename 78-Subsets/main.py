class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        res = []
        
        def dfs(i: int, path: List[int]) -> None:
            if i == n:
                res.append(path)
                return
            
            dfs(i+1, path)
            dfs(i+1, path + [nums[i]])
        
        dfs(0, [])
        return res
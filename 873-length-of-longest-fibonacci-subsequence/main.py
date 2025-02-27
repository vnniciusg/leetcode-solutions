class Solution:
    def lenLongestFibSubseq(self, arr: List[int]) -> int:
        
        idx = {x: i for i, x in enumerate(arr)}
        dp = {}
        max_len = 0
        n = len(arr)


        for k in range(n):
            for j in range(k):

                i_val = arr[k] - arr[j]

                if i_val < arr[j] and i_val in idx:
                    i = idx[i_val]
                    dp[j, k] = dp.get((i, j), 2) + 1
                    max_len = max(max_len, dp[j, k])


        return max_len if max_len >= 3 else 0
    

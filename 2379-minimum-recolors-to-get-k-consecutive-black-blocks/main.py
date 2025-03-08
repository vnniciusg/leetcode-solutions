class Solution:
    def minimumRecolors(self, blocks: str, k: int) -> int:
        ans = k
        c = 0
        n = len(blocks)
        for i in range(n):
            c += blocks[i] == 'B'
            c -= i - k >= 0 and blocks[i - k] == 'B'
            ans = min(ans, k - c)
        return ans

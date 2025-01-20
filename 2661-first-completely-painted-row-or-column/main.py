class Solution:
    def firstCompleteIndex(self, arr: List[int], mat: List[List[int]]) -> int:
        m = len(mat)
        n = len(mat[0]) if m > 0 else 0
        pos = {}
        for i in range(m):
            for j in range(n):
                val = mat[i][j]
                pos[val] = (i, j)

        
        rows, cols = [0] * m, [0] * n
        for idx, val in enumerate(arr):
            i, j = pos[val]
            rows[i] += 1
            if rows[i] == n: return idx
            cols[j] += 1
            if cols[j] == m: return idx

        return -1

from typing import List

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:

        if not matrix:
            return []

        m, n = len(matrix), len(matrix[0])
        res = []
        left, right, top, bottom = 0, n - 1, 0, m - 1

        while left <= right and top <= bottom:

            for j in range(left, right+1):
                res.append(matrix[top][j])
            top+=1

            for i in range(top, bottom+1):
                res.append(matrix[i][right])
            right -= 1

            if top <= bottom:
                for j in range(right, left - 1, -1):
                    res.append(matrix[bottom][j])
                bottom -= 1
            
            if left <= right:
                for i in range(bottom, top - 1, -1):
                    res.append(matrix[i][left])
                left += 1

        
        return res
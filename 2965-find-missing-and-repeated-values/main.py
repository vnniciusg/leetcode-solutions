class Solution:
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        n = len(grid)
        count = Counter(num for row in grid for num in row)

        a = b = -1
        for num in range(1, n * n + 1):
            if count[num] == 2:
                a = num
            elif num not in count:
                b = num 

        return [a, b]

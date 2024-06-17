from typing import List

class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        
        n = len(arr)

        count = 0
        for i in range(n):

            if arr[i] & 1 != 0:
                count+=1
            else:
                count = 0

            if count == 3:
                return True
        

        return False
        
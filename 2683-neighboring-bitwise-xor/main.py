class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        xor = 0
        for element in derived: 
            xor ^= element
        return xor == 0

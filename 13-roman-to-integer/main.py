class Solution:
    def romanToInt(self, s: str) -> int:
        roman_dict = {
            'I':1,
            'V':5,
            'X':10,
            'L':50,
            'C':100,
            'D':500,
            'M':1000,
        }

        total, prev = 0, 0

        for char in s[::-1]:
            value = roman_dict.get(char)

            if value >= prev:
                total += value
            else:
                total -= value
            
            prev = value
        
        return total

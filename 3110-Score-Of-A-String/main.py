class Solution:
    def scoreOfString(self, s: str) -> int:
        
        n = len(s)
        i, score = 0, 0
        for i in range(n - 1):
            score += abs(self.parse_letter_to_number(s[i]) -  self.parse_letter_to_number(s[i+1]))
        
        return score
    
    def parse_letter_to_number(self, letter:str) -> int:
        return ord(letter.lower()) - ord('a') + 1



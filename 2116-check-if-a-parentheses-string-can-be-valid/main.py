class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        n = len(s)
        
        if n % 2: return False

        free = balance = 0
        for i in range(n):
            if locked[i] == '1': balance += 1 if s[i] == '(' else -1
            else: free += 1

            if balance + free < 0: return False

        free = balance = 0
        for i in range(n-1, -1, -1):
            if locked[i] == '1': balance += 1 if s[i] == ')' else -1
            else: free += 1

            if balance + free < 0: return False

        return True

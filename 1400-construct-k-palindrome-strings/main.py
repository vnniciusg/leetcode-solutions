class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        
        if len(s) < k: return False
        if len(s) == k: return True

        freq = [s.count(chr(i + ord('a'))) for i in range(26)]
        odd_count = 0

        for count in freq:
            if count & 1: odd_count += 1

        return odd_count <= k

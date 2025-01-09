class Solution:
    def prefixCount(self, words: List[str], pref: str) -> int:
        
        count: int = 0
        
        for word in words:
            if word.startswith(pref):
                count += 1

        return count

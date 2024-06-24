class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        
        if needle not in haystack or len(haystack) < len(needle):
            return -1
        
        n = len(haystack)
        for i in range(n):
            if haystack[i: i + len(needle)] == needle:
                return i

        return -1


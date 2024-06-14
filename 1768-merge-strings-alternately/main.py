from itertools import zip_longest

class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        ans = ''
        for char1, char2 in zip_longest(word1, word2, fillvalue=''):
            ans += char1 + char2
        return ans

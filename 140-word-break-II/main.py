from typing import List

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:

        word_set = set(wordDict)
        def backtrack(start, path):
            if start == len(s):
                sentences.append(' '.join(path))
                return
            
            for end in range(start + 1, len(s) + 1):
                if s[start:end] in word_set:
                    backtrack(end, path + [s[start:end]])

        sentences = []
        backtrack(0, [])
        return sentences
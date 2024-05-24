
from collections import Counter
from typing import List

class Solution:
    def maxScoreWords(self, words: List[str], letters: List[str], score: List[int]) -> int:

        letter_count = Counter(letters)
        points = {word: sum([score[ord(char) - ord('a')] for char in word]) for word in words }

        def dfs(i: int) -> int:
            if i == len(words):
                return 0

            word = words[i]
            count = Counter(word)
            if all([letter_count[char] >= count[char] for char in count]):
                for char in count:
                    letter_count[char] -= count[char]
                score = points[word] + dfs(i + 1)
                for char in count:
                    letter_count[char] += count[char]
                
                return max(score, dfs(i + 1))

            return dfs(i+1)

        return dfs(0)

    


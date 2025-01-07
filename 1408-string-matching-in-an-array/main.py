class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:

        s = ' '.join(words)
        res = []

        for word in words:
            if s.count(word) > 1:
                res.append(word)
        return res 

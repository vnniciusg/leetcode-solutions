class Solution:
    def commonChars(self, words: List[str]) -> List[str]:
        
        if not words or len(words) == 1:
            return list(words[0]) if words else []
        

        common_freq = Counter(words[0])

        for word in words[1:]:
            word_freq = Counter(word)
            common_freq &= word_freq
        
        ans = []
        for char, freq in common_freq.items():
            ans.extend([char] * freq)
        

        return ans
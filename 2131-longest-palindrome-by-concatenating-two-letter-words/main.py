class Solution:
    def longestPalindrome(self, words: List[str]) -> int:

        freq = defaultdict(int)
        ans = f = 0

        for word in words:
            freq[word] += 1

        
        for k, v in freq.items():
            x, y = k
            if x == y:
                if v & 1 and not f: f = 2
                ans += 2 * (v//2)
            elif y+x in freq:
                ans += min(freq[k], freq[y+x])

        return 2 * ans + f        

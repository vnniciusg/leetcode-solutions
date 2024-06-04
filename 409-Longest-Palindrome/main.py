class Solution:
    def longestPalindrome(self, s: str) -> int:
        
        freq = defaultdict(int)

        for char in s:
            freq[char] = freq.get(char, 0) + 1
        
        lenght = 0
        odd_found = False

        for count in freq.values():
            if count & 1 == 0:
                lenght += count
            else:
                lenght += count - 1
                odd_found = True
        
        if odd_found:
            lenght += 1
        
        return lenght

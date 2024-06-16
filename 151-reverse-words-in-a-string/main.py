class Solution:
    def reverseWords(self, s: str) -> str:
        
        l = s.split()
        n = len(l)
        for i in range(n// 2):
            l[i], l[n - i - 1] = l[n - i - 1], l[i]

        return " ".join(l)

        

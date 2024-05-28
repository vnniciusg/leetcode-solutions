class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:

        n = len(s)
        diff = [abs(ord(s[i]) - ord(t[i])) for i in range(n)]

        max_lenght = 0
        curr_cost = 0
        start = 0


        for end in range(n):
            curr_cost += diff[end]
            
            while curr_cost > maxCost:
                curr_cost -= diff[start]
                start += 1
            
            max_lenght = max(max_lenght, end - start + 1)


        return max_lenght



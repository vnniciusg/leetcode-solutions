class Solution:
    def shiftingLetters(self, s: str, shifts: List[List[int]]) -> str:
        
        n = len(s)
        diff = [0] * (n + 1)
        
        for start, end, direction in shifts:
            diff[start] += 1 if direction == 1 else -1
            diff[end + 1] -= 1 if direction == 1 else -1
        
        curr_shift = 0
        ans = list(s)
        for i in range(n):
            curr_shift = (curr_shift + diff[i]) % 26
            curr_chr = ord(ans[i]) - ord('a')

            new_chr = (curr_chr + curr_shift) % 26
            ans[i] = chr(new_chr + ord('a'))

        return "".join(ans)

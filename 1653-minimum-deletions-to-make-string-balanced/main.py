class Solution:
    def minimumDeletions(self, s: str) -> int:
        b_cnt, num_deletions = 0, 0

        for char in s:
            if char == 'b': b_cnt += 1
            else:
                num_deletions = min(num_deletions + 1, b_cnt)
            
        return num_deletions

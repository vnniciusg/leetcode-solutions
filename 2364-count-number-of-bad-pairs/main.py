class Solution:
    def countBadPairs(self, nums: List[int]) -> int:
        
        n = len(nums)
        total_pairs = n * (n -1)// 2
        good_pairs = 0

        diff_count = defaultdict(int)

        for i, num in enumerate(nums):
            diff = num - i
            good_pairs += diff_count[diff]
            diff_count[diff] += 1

        return total_pairs - good_pairs

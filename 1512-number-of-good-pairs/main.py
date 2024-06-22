class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        
        count_values = defaultdict(int)

        count_good_pairs = 0
        for num in nums:
            count_good_pairs += count_values[num]
            count_values[num] += 1

        return count_good_pairs

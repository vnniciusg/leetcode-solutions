class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        
        count_pairs = 0
        count_map = {}

        for num in nums:
            complement = k - num
            if count_map.get(complement, 0) > 0:
                count_pairs += 1
                count_map[complement] -= 1
            else:
                count_map[num] = count_map.get(num, 0) + 1
        
        return count_pairs

class Solution:
    def tupleSameProduct(self, nums: List[int]) -> int:
        
        freq = defaultdict(int)
        n = len(nums)
        for i in range(n - 1):
            for j in range(i + 1, n):
                freq[nums[i] * nums[j]] += 1

        count = 0
        for product, frequency in freq.items():
            count += 8 * (frequency * (frequency-1)) / 2

        return int(count)

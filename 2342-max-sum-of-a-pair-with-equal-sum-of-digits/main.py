class Solution:
    def maximumSum(self, nums: List[int]) -> int:
        
        groups = defaultdict(list)
        for num in nums:
            digit_sum = sum([int(char) for char in str(num)])
            groups[digit_sum].append(num)

        max_sum = -1

        for group in groups.values():
            if len(group) < 2: continue
            group.sort(reverse=True)
            max_sum = max(max_sum, group[0] + group[1])

        return max_sum

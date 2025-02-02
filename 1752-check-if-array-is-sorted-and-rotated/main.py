class Solution:
    def check(self, nums: List[int]) -> bool:
        
        n = len(nums)
        sorted_nums = sorted(nums)
        if nums == sorted_nums:
            return True

        for i in range(n):
            is_match = True

            for j in range(n):
                if sorted_nums[j] != nums[(i + j) % n]:
                    is_match = False
                    break

            if is_match:
                return True

        return False

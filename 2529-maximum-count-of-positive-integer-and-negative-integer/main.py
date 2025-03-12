class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        
        def binary_search(target):
            l, r = 0, len(nums) - 1
            result = len(nums)

            while l <= r:
                mid = l + (r - l) // 2
                if nums[mid] < target:
                    l = mid + 1
                else:
                    r = mid - 1
                    result = mid

            return result

        return max(binary_search(0), len(nums) - binary_search(1))

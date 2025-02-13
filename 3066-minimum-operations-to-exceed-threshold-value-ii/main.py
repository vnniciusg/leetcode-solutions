class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        
        def _helper(nums: List[int], n_operations: int) -> int:
            
            heapq.heapify(nums)

            while nums[0] < k:

                if len(nums) < 2:
                    return n_operations
                
                smallest = heapq.heappop(nums)
                second_smallest = heapq.heappop(nums)

                new_element = smallest * 2 + second_smallest
                heapq.heappush(nums, new_element)

                n_operations += 1
            
            return n_operations

        return _helper(nums, 0)

from typing import List
from collections import defaultdict
"""

You are given a 0-indexed integer array nums and an integer k.

A subarray is called equal if all of its elements are equal. Note that the empty subarray is an equal subarray.

Return the length of the longest possible equal subarray after deleting at most k elements from nums.

A subarray is a contiguous, possibly empty sequence of elements within an array.

For each number x in nums, create a sorted list indicesx of all indices i such that nums[i] == x.
On every indicesx, execute a sliding window technique.
For each indicesx, find i, j such that (indicesx[j] - indicesx[i]) - (j - i) <= k and j - i + 1 is maximized.
The answer would be the maximum of j - i + 1 for all indicesx.

Example 1:

Input: nums = [1,3,2,3,1,3], k = 3
Output: 3
Explanation: It's optimal to delete the elements at index 2 and index 4.
After deleting them, nums becomes equal to [1, 3, 3, 3].
The longest equal subarray starts at i = 1 and ends at j = 3 with length equal to 3.
It can be proven that no longer equal subarrays can be creat
"""

class Solution:
    def longestEqualSubarray(self, nums: List[int], k: int) -> int:

        n = len(nums)
        left, res = 0, 0

        freq = defaultdict(int)
        cnt = 0

        for right in range(n):
            freq[nums[right]] += 1

            cnt = max(cnt, freq[nums[right]])

            if cnt + k < right - left + 1:
                freq[nums[left]] -= 1
                left +=1
            else:
                res = max(res,  cnt)

        return res
    

s = Solution()

print(s.longestEqualSubarray([1,3,2,3,1,3], 3), "solution 1")
print(s.longestEqualSubarray([1,1,2,2,1,1], 2), "solution 2")

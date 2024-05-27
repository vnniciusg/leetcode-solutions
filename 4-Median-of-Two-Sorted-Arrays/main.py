class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:

        arr = sorted(nums1 + nums2)

        mid = len(arr)// 2

        if len(arr) & 1 != 0:
            return arr[mid]

        return (arr[mid] + arr[mid-1]) / 2 
        

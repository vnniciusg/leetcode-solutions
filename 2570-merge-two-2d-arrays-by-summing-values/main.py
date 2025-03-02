class Solution:
    def mergeArrays(self, nums1: List[List[int]], nums2: List[List[int]]) -> List[List[int]]:
        
        _dict = defaultdict(int)

        for _id, _val in nums1:
            _dict[_id] = _val

        for _id, _val in nums2:
            _dict[_id] = _dict.get(_id, 0) + _val

        ans = []
        for key, value in sorted(_dict.items()):
            ans.append([key, value])

        return ans

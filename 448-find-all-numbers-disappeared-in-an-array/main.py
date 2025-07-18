class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        
        _counter = Counter(nums)
        _numbers = [i + 1 for i in range(len(nums))]
        
        _missing = []
        for _number in _numbers:
            if _number not in _counter.keys():
                _missing.append(_number)

        return _missing

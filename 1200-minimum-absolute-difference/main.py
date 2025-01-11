class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        _min = float('inf')

        n = len(arr) -1
        for i in range(n):
            _min = min(_min, abs(arr[i] - arr[i+1]))
        
        ans = []
        for i in range(n):
            if abs(arr[i] - arr[i+1]) == _min:
                ans.append([arr[i], arr[i+1]])

        return ans

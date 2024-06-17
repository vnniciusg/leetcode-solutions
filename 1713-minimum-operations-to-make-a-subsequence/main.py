class Solution:
    def minOperations(self, target: List[int], arr: List[int]) -> int:
       
        target_index = {num: i for i, num in enumerate(target)}
        target_indices = [target_index[num] for num in arr if num in target_index]

        lis = []

        for index in target_indices:
            pos = bisect.bisect_left(lis, index)
            if pos == len(lis):
                lis.append(index)
            else:
                lis[pos] = index
        

        return len(target) - len(lis)


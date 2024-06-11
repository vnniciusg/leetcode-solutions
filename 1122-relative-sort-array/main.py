class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        
        order_map = {num: idx for idx, num in enumerate(arr2)}

        count_map = {}
        for num in arr1:
            count_map[num] = count_map.get(num, 0) + 1


        ans = [] 
        for num in arr2:
            if num in count_map:
                ans.extend([num] * count_map[num])
                del count_map[num]
        

        remaining = sorted([num for num in count_map for _ in range(count_map[num])])
        ans.extend(remaining)


        return ans


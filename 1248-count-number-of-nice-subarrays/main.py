class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:

        
        count = 0
        odd_count = 0
        hash_map = {0: 1}

        for num in nums:
            if num & 1 != 0:
                odd_count += 1
                
            if odd_count - k in hash_map:
                count += hash_map[odd_count - k]

            if odd_count in hash_map:
                hash_map[odd_count] += 1
            else:
                hash_map[odd_count] = 1

        return count



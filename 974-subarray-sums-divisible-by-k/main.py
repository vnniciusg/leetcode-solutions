class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        
        count, prefix_sum = 0, 0
        mod_count = {0: 1}

        for num in nums:
            prefix_sum += num
            mod = prefix_sum % k

            if mod < 0:
                mod += k
            
            count += mod_count.get(mod, 0)
            mod_count[mod] = mod_count.get(mod, 0) + 1
      

        return count

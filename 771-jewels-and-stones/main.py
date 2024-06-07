class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        
        s = {stone: stones.count(stone) for stone in set(stones)}
        j = [jewel for jewel in jewels]
        
        count = 0
        for jewel in j:
            if jewel in s:
                count += s[jewel]
        

        return count

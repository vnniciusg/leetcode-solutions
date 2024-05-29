class Solution:
    def numSteps(self, s: str) -> int:

        n = int(s, 2)

        count_steps = 0
        while n != 1:
            if n & 1 != 0:
                n += 1
            else:
                n //= 2
                
            count_steps+=1
        
        return count_steps
        

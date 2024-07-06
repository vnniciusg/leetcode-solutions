class Solution:
    def passThePillow(self, n: int, time: int) -> int:
        
        position, direction = 1, 1

        for _ in range(time):
            position += direction
            if position == n:
                direction = -1
            elif position == 1:
                direction = 1
        
        return position

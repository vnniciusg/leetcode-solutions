class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        
        friends = [i for i in range(1, n + 1)]
        idx = 0

        while len(friends) > 1:
            idx = (idx + k - 1) % len(friends) 
            friends.pop(idx)

        return friends[0]
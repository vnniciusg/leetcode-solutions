class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        
        a, b = edges[0]
        
        if a in edges[1]:
            return a
        
        return b

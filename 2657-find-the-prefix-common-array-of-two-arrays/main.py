class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
                
        n = len(A)
        ans = []
        set_a, set_b = set(), set()

        for idx in range(n):
            set_a.add(A[idx])
            set_b.add(B[idx])

            common = len(set_a.intersection(set_b))
            ans.append(common)
        
        return ans

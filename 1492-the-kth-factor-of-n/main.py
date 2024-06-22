class Solution:
    def kthFactor(self, n: int, k: int) -> int:
        
        factor_list = [idx for idx in range(1, n + 1) if n % idx == 0]

        if k - 1 >= len(factor_list):
            return -1

        return factor_list[k - 1]

class Solution:
    def minimizeXor(self, num1: int, num2: int) -> int:

        k = bin(num2).count('1')
        c = bin(num1).count('1')

        if c <= k:
            x = num1
            remaining = k - c
            for bit in range(0, 32):
                if not(x & (1 << bit)) and remaining > 0:
                    x |= (1 << bit)
                    remaining -= 1
        
        else:
            set_bits = [bit for bit in range(31, -1, -1) if num1 & (1 << bit)]
            top_k_bits = set_bits[:k]
            x = 0
            for bit in top_k_bits:
                x |= (1 << bit)
        
        return x
        

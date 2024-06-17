class Solution:
    def countTriplets(self, arr: List[int]) -> int:
        
        n = len(arr)
        prefix = [0] * (n + 1)

        for i in range(n):
            prefix[i+1] = prefix[i] ^ arr[i]

        count = 0
        for j in range(n):
            for k in range(j + 1, n):
                if prefix[j] == prefix[k + 1]:
                    count += (k - j)

        return count

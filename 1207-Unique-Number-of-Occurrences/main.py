class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        
        freq = {num: arr.count(num) for num in arr}
        freq_values = set(freq.values())
        
        return len(freq_values) == len(freq)


class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        
        freq = defaultdict(int)

        for num in arr:
            freq[num] += 1

        occurrence_set = set(freq.values())
        return len(occurrence_set) == len(freq)



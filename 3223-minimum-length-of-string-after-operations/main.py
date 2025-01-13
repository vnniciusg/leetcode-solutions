class Solution:
    def minimumLength(self, s: str) -> int:
        freq = Counter(s)

        del_count = 0
        for frequency in freq.values():
            if frequency & 1: 
                del_count += frequency - 1
            else:
                del_count += frequency - 2
            
        return len(s) - del_count

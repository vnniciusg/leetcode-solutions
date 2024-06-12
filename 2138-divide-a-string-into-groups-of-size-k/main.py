class Solution:
    def divideString(self, s: str, k: int, fill: str) -> List[str]:
        
        split_string_list = [s[i:i+k] for i in range(0, len(s), k)]

        if len(split_string_list[-1]) != k:
            split_string_list[-1] += fill * (k - len(split_string_list[-1]))

        return split_string_list

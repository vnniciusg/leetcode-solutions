class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        
        vowels = set(['a', 'e', 'i', 'o', 'u'])
        n = len(words)


        prefix_sum = [0] * n
        _sum = 0
        for i in range(n):
            curr_word = words[i]
            if curr_word[0] in vowels and curr_word[-1] in vowels:
                _sum += 1
            
            prefix_sum[i] = _sum
 

        m = len(queries)
        ans = [0] * m
        for i in range(m):
            curr_query = queries[i]
            ans[i] = prefix_sum[curr_query[1]] - (0 if curr_query[0] == 0 else prefix_sum[curr_query[0] - 1])

        return ans

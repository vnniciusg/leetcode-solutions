class Solution:
    def countOfSubstrings(self, word: str, k: int) -> int:
        
        def at_least_k(k: int):
            vowel_count = defaultdict(int)
            non_vowel_count = 0
            left = 0
            result = 0
            n = len(word)

            for right in range(n):
                if word[right] in "aeiou":
                    vowel_count[word[right]] += 1
                else:
                    non_vowel_count += 1

                while len(vowel_count) == 5 and non_vowel_count >= k:
                    result += (n - right)
                    if word[left] in "aeiou":
                        vowel_count[word[left]] -= 1
                        if vowel_count[word[left]] == 0:
                            del vowel_count[word[left]]
                    else:
                        non_vowel_count -= 1
                    left += 1
            return result
        
        return at_least_k(k) - at_least_k(k + 1)

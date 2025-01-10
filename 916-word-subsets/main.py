class Solution:
    def wordSubsets(self, words1: List[str], words2: List[str]) -> List[str]:
        
        counter_sub = Counter()
        for word in words2:
            counter_sub |= Counter(word)
        return [word for word in words1 if Counter(word)>= counter_sub]

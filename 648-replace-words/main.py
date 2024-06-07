class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:

        root_set = set(dictionary)
        
        def findRoot(word: str) -> str:
            for i in range(1, len(word) + 1):
                prefix = word[:i]
                if prefix in root_set:
                    return prefix
            return word
        
        words = sentence.split()
        replaced_words = map(findRoot, words)

        return ' '.join(replaced_words)

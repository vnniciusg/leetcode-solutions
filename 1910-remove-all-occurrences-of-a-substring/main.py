class Solution:
    def removeOccurrences(self, s: str, part: str) -> str:
        
        def _helper(substr: str) -> str:
            
            if part not in substr:
                return substr

            substr = substr.replace(part, '', 1)
            substr = _helper(substr)
            return substr


        return _helper(s)

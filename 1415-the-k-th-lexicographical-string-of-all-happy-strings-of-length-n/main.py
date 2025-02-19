class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        
        def backtrack(curr: str, count: int = 0) -> tuple[bool, int, str]:

            if len(curr) == n:
                count += 1
                if count == k: 
                    return True, count, curr
                
                return False, count, ""

            for char in 'abc':
                if not curr or curr[-1] != char:
                    found, count, ans = backtrack(curr + char, count)
                    if found:
                        return True, count, ans
                
            return False, count, ""

        found, _, answer = backtrack("", 0)
        return answer if found else ""           


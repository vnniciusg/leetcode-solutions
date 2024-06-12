class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        
        def is_divisible_by(v: int, div: int) -> bool:
            if v % div != 0:
                return False
                
            return True

        ans = []
        for i in range(1, n + 1):
            if is_divisible_by(i, 15):
                ans.append("FizzBuzz")
            elif is_divisible_by(i, 3):
                ans.append("Fizz")
            elif is_divisible_by(i, 5):
                ans.append("Buzz")
            else:
                ans.append(str(i))

        return ans

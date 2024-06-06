class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        return [int(num) for num in str(int("".join(str(digit) for digit in digits)) + 1)]

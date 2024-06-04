class Solution:
    def largestPalindromic(self, num: str) -> str:
        count = Counter(num)
        first_half, middle = [], ""

        for digit in sorted(count.keys(), reverse=True):
            if count[digit] & 1 != 0:
                middle = max(middle, digit)
            first_half.append(digit * (count[digit] // 2))
        
        first_half_str = "".join(first_half)
        
        first_half_str = first_half_str.lstrip('0')
        if first_half_str == "":
            return middle if middle else "0"
        
        palindrome = first_half_str + middle + first_half_str[::-1]
        return palindrome
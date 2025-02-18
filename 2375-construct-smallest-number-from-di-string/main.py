class Solution:
    def smallestNumber(self, pattern: str) -> str:
        ans = []
        num_stack = []
        n = len(pattern)

        for idx in range(n + 1):

            num_stack.append(idx + 1)

            if idx == n or pattern[idx] == "I":
                while num_stack:
                    ans.append(str(num_stack.pop()))
                    print(ans)

        return "".join(ans)

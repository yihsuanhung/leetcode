class Solution:
    def removeStars(self, s: str) -> str:
        stack = []
        for c in list(s):
            if c is not "*":
                stack.append(c)
            else:
                stack.pop()
        return "".join(stack)

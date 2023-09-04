class Solution:
    def reverseWords(self, s: str) -> str:
        # strip
        result = s.strip()
        # split
        result = result.split(" ")
        # cleanup
        result = [i for i in result if i != ""]
        # reverse
        result.reverse()

        return " ".join(result)

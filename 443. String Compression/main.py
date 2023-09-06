class Solution:
    def compress(self, chars: List[str]) -> int:
        i = 0
        result = 0
        while i < len(chars):
            group_len = 1
            while (i+group_len < len(chars)) and (chars[i+group_len] == chars[i]):
                group_len += 1
            chars[result] = chars[i]
            result += 1
            if group_len > 1:
                s = str(group_len)
                chars[result:result+len(s)] = list(s)
                result += len(s)
            i += group_len
        return result

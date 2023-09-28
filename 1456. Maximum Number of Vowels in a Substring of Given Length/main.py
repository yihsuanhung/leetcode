class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        chars = list(s)
        vowels = set(["a", "e", "i", "o", "u"])
        max_n_vow = 0
        cur_n_vow = 0
        start = 0
        for end in range(len(chars)):
            if chars[end] in vowels:
                cur_n_vow += 1
            if end - start + 1 > k:
                if chars[start] in vowels:
                    cur_n_vow -= 1
                start += 1
            max_n_vow = max(max_n_vow, cur_n_vow)
        return max_n_vow if k <= len(chars) else 0

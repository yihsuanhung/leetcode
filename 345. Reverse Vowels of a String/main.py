class Solution:
    def reverseVowels(self, s: str) -> str:

        vowels = "aeiouAEIOU"
        word = list(s)

        # two pointer
        left = 0
        right = len(s)-1

        while left < right:
            # find the vowel from the left
            # if use "or", it will change back
            while left < right and vowels.find(word[left]) == -1:
                left += 1
            # find the vowel from the right
            while left < right and vowels.find(word[right]) == -1:
                right -= 1

            word[left], word[right] = word[right], word[left]
            left += 1
            right -= 1

        return "".join(word)

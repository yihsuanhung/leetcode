class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        char_set_1, char_set_2 = set(word1), set(word2)
        freq_map_1, freq_map_2 = {}, {}

        for c1 in word1:
            freq_map_1[c1] = freq_map_1.get(c1, 0) + 1

        for c2 in word2:
            freq_map_2[c2] = freq_map_2.get(c2, 0) + 1

        return char_set_1 == char_set_2 and sorted(freq_map_1.values()) == sorted(freq_map_2.values())

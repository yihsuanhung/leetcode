class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        num_map = {}
        o_set = set()
        for i in arr:
            if i in num_map:
                num_map[i] += 1
            else:
                num_map[i] = 1
        return len(num_map.values()) == len(set(num_map.values()))

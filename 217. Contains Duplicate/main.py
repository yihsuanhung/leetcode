from typing import List


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        number_set = set()

        for n in nums:
            if n not in number_set:
                number_set.add(n)
            else:
                return True
        return False


solution = Solution()
answer = solution.containsDuplicate([1, 2, 3])
print(answer)

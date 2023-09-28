class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        start, using, longest = 0, 0, 0
        # 展開窗戶
        for end in range(len(nums)):
            if nums[end] == 0:
                using += 1
                # 縮小左指針，縮小窗戶
            while using > 1:
                if nums[start] == 0:
                    using -= 1
                start += 1
            longest = max(longest, end - start + 1)
        return longest - 1

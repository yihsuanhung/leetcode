class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        s = 0
        max_avg = float("-inf")
        curr_sum = 0
        for e in range(len(nums)):
            # keep caculate sum untile end
            curr_sum += nums[e]
            # run this code after e reach the end of the window
            if e >= k-1:
                curr_avg = curr_sum / k
                max_avg = max(curr_avg, max_avg)
                curr_sum -= nums[s]
                s += 1
        return max_avg

class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        nums1_set = set(nums1)
        nums2_set = set(nums2)
        answer = [[], []]
        for n1 in nums1_set:
            if n1 not in nums2_set:
                answer[0].append(n1)
        for n2 in nums2_set:
            if n2 not in nums1_set:
                answer[1].append(n2)
        return answer

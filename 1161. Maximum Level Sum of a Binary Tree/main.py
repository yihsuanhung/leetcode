class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        max_sum = float("-inf")
        level, max_level = 0, 0

        q = []
        q.append(root)

        while q:
            level += 1
            level_size = len(q)
            level_sum = 0
            for i in range(level_size):
                top = q.pop(0)
                level_sum += top.val
                if top.left:
                    q.append(top.left)
                if top.right:
                    q.append(top.right)
            if level_sum > max_sum:
                max_level = level
                max_sum = level_sum

        return max_level

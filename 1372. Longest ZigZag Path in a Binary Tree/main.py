class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:

        max_count = 0

        def dfs(node, prv_dir, count):
            if not node:
                return 0

            left_count, right_count = 0, 0

            if prv_dir == 'left':
                left_count = dfs(node.left, 'left', 1)
                right_count = dfs(node.right, 'right', count+1)
            elif prv_dir == 'right':
                left_count = dfs(node.left, 'left', count+1)
                right_count = dfs(node.right, 'right', 1)
            else:
                left_count = dfs(node.left, 'left', 1)
                right_count = dfs(node.right, 'right', 1)

            return max(left_count, right_count, count)

        return dfs(root, '', 0)

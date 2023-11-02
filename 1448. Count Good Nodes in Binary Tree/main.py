class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        def dfs(node, max_value):
            if not node:
                return 0
            good_count = 1 if node.val >= max_value else 0
            max_value = max(max_value, node.val)
            good_count += dfs(node.left, max_value)
            good_count += dfs(node.right, max_value)
            return good_count
        return dfs(root, root.val)

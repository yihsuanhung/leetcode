class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        return self.dfs(root, 0)

    def dfs(self, root, depth):
        if root is None:
            return depth
        return max(self.dfs(root.left, depth+1), self.dfs(root.right, depth+1))

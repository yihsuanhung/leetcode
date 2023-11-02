# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        l1, l2 = [], []
        self.dfs(root1, l1)
        self.dfs(root2, l2)
        return l1 == l2

    def dfs(self, root, leaves):
        if root is None:
            return
        if root.left is None and root.right is None:
            leaves.append(root.val)
            return
        self.dfs(root.left, leaves)
        self.dfs(root.right, leaves)

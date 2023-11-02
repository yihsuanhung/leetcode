# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if not root:
            return root

        def findLargest(node):
            while node.right:
                node = node.right
            return node

        def delete(node, target):
            # in leaf, not found
            if not node:
                return node
            # node.right will attach to a new subtree
            if target > node.val:
                node.right = delete(node.right, target)
            # node.left will attach to a new subtree
            elif target < node.val:
                node.left = delete(node.left, target)
            # found
            else:
                # case: target is in leaf
                if not node.left and not node.right:
                    return None
                # case: not leaf, only right
                elif not node.left:
                    return node.right
                # case: not leaf, only left
                elif not node.right:
                    return node.left
                # case: not leaf, with left and right
                else:
                    left_largest = findLargest(node.left)
                    node.val = left_largest.val
                    node.left = delete(node.left, left_largest.val)
            return node

        return delete(root, key)

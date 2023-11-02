class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:

        def dfs(node, path):
            if not node:
                return 0

            path.append(node.val)
            path_sum, path_count = 0, 0

            for i in range(len(path)-1, -1, -1):
                path_sum += path[i]
                if path_sum == targetSum:
                    path_count += 1

            path_count += dfs(node.left, path)
            path_count += dfs(node.right, path)

            path.pop()

            return path_count

        return dfs(root, [])

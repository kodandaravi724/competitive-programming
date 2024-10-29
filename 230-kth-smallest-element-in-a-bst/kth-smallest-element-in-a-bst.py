# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def traverse(self, root: Optional[TreeNode], k:int, current: int) -> int:
        if root is not None:
            left_value = self.traverse(root.left, k, current)
            current[0]+=1
            currentElementOrder = current[0]
            right_value = self.traverse(root.right, k, current)
            if currentElementOrder == k:
                return root.val
            if currentElementOrder != k:
                return max(left_value, right_value)
        else:
            return -1



    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        current = [0]
        return self.traverse(root, k, current)
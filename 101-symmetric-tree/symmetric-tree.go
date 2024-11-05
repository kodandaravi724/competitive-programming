/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkIsSymmetric(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil && right != nil {
        return false
    }
    if left != nil && right == nil {
        return false
    }
    if left.Val != right.Val {
        return false
    }
    return checkIsSymmetric(left.Left, right.Right) && checkIsSymmetric(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
    return checkIsSymmetric(root.Left, root.Right)
} 
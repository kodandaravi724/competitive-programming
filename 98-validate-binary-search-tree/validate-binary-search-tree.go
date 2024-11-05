
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkIfValid(root *TreeNode, leftBound int, rightBound int) bool{
    if root == nil {
        return true
    }
    if root.Val > leftBound && root.Val < rightBound {
        return checkIfValid(root.Left, leftBound, root.Val) && checkIfValid(root.Right, root.Val, rightBound)
    } else {
        return false
    }
}

func isValidBST(root *TreeNode) bool {
    return checkIfValid(root, -100000000000, 100000000000)
}
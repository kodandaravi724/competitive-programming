import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkIfValid(root *TreeNode, leftBound float64, rightBound float64) bool{
    if root == nil {
        return true
    }
    if float64(root.Val) > leftBound && float64(root.Val) < rightBound {
        return checkIfValid(root.Left, leftBound, float64(root.Val)) && checkIfValid(root.Right, float64(root.Val), rightBound)
    } else {
        return false
    }
}

func isValidBST(root *TreeNode) bool {
    leftBound := math.Inf(-1)
    rightBound := math.Inf(1)
    return checkIfValid(root, leftBound, rightBound)
}
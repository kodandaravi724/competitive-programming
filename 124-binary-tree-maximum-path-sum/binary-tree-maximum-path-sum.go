/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, maxSum *int) int{
     if root == nil {
        return 0
    }
    leftPathSum := max(0, traverse(root.Left, maxSum))
    rightPathSum := max(0 ,traverse(root.Right, maxSum))
    *maxSum = max(*maxSum, leftPathSum + root.Val + rightPathSum)
    if leftPathSum > rightPathSum {
        return leftPathSum + root.Val
    } else {
        return rightPathSum + root.Val
    }
}

func maxPathSum(root *TreeNode) int {
    var maxSum int = -1001
    traverse(root, &maxSum)
    return maxSum
}
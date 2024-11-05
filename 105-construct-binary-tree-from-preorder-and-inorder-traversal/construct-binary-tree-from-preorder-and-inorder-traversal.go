/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getIndex(arr []int, element int) int {
    for index, arr_element := range(arr) {
        if arr_element == element {
            return index
        }
    }
    return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    rootVal := preorder[0]
    root := new(TreeNode)
    root.Val = rootVal
    index := getIndex(inorder, rootVal)
    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
    return root
}
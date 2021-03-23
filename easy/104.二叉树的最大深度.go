/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDeep:=maxDepth(root.Left) + 1
    rightDeep:=maxDepth(root.Right) + 1
    if leftDeep >= rightDeep {
        return leftDeep
    }
    return rightDeep
}
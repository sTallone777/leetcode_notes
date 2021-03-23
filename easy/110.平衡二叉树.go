/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root != nil {
        m := getHeight(root.Left) - getHeight(root.Right)
        if m >= -1 && m <= 1{
            if isBalanced(root.Left) && isBalanced(root.Right){
                return true
            }
        }
        return false
    }
    return true 
}

func getHeight(node *TreeNode) int {
    if node != nil {
        leftHheight := getHeight(node.Left)+1
        rightHeight := getHeight(node.Right)+1
        if leftHheight >= rightHeight {
            return leftHheight
        }
        return rightHeight
    }
    return 0
}
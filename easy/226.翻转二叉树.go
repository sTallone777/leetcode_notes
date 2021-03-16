/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		n := invertTree(root.Left)
		root.Left = invertTree(root.Right)
		root.Right = n
		return root
	}
	return nil
}
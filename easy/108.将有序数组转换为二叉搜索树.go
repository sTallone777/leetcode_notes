/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) > 0{
        var node *TreeNode = new(TreeNode)
        halfNum:= len(nums) / 2
        node.Val = nums[halfNum]
        node.Left = sortedArrayToBST(nums[:halfNum])
        node.Right = sortedArrayToBST(nums[halfNum+1:])
        return node
    }
    return nil
}
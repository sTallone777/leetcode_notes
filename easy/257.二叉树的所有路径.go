/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	traversal(root, "")
	return paths
}

func traversal(root *TreeNode, path string) {
	if root != nil {
		tmp := path
		tmp += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, tmp)
		} else {
			tmp += "->"
			traversal(root.Left, tmp)
			traversal(root.Right, tmp)
		}
	}
}

// @lc code=end


package main

//给定一个二叉树，返回它的 前序 遍历。
//binary-tree-preorder-traversa
var preorderTraversalRes []int

func preorderTraversal(root *TreeNode) []int {
	preorderTraversalRes = []int{}
	PreorderTraversal(root)
	return preorderTraversalRes
}

func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	} else {
		preorderTraversalRes = append(preorderTraversalRes, root.Val)
		PreorderTraversal(root.Left)
		PreorderTraversal(root.Right)
	}
	return
}

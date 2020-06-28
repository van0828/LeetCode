package main

//给定一个二叉树，返回它的中序 遍历。
//binary-tree-inorder-traversal
var inorderTraversalRes []int

func inorderTraversal(root *TreeNode) []int {
	inorderTraversalRes = []int{}
	InorderTraversal(root)
	return inorderTraversalRes
}

func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	} else {
		InorderTraversal(root.Left)
		inorderTraversalRes = append(inorderTraversalRes, root.Val)
		InorderTraversal(root.Right)
	}
	return
}

package main

import "math"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//validate-binary-search-tre

//var inorderTraversalRes []int
//
//func inorderTraversal(root *TreeNode) []int {
//	inorderTraversalRes = []int{}
//	InorderTraversal(root)
//	return inorderTraversalRes
//}
//
//func InorderTraversal(root *TreeNode) {
//	if root == nil {
//		return
//	} else {
//		InorderTraversal(root.Left)
//		inorderTraversalRes = append(inorderTraversalRes, root.Val)
//		InorderTraversal(root.Right)
//	}
//	return
//}

//func isValidBST(root *TreeNode) bool {
//	inorderTraversal(root)
//	for i := 0; i < len(inorderTraversalRes)-1; i++ {
//		if inorderTraversalRes[i] >= inorderTraversalRes[i+1] {
//			return false
//		}
//	}
//	return true
//}

func isValidBST(root *TreeNode) bool {
	var IsValidBST func(root *TreeNode) bool
	last := math.MinInt64
	IsValidBST = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if IsValidBST(root.Left) {
			if last < root.Val {
				last = root.Val
				return IsValidBST(root.Right)
			}
		}
		return false
	}
	return IsValidBST(root)
}

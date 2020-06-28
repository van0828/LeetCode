package main

//给定一个二叉树，原地将它展开为链表。
//
//例如，给定二叉树
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
//flatten-binary-tree-to-linked-list

// 后序遍历防止丢失左孩子
func flatten(root *TreeNode)  {
	var handler func(root *TreeNode)
	var prev *TreeNode
	handler = func(root *TreeNode) {
		if root == nil {
			return
		}
		handler(root.Right)
		handler(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}
	handler(root)
}
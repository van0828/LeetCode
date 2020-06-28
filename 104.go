package main

import (
	"github.com/eapache/queue"
)


func main() {

}

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//maximum-depth-of-binary-tree
func maxDepth(root *TreeNode) int {
	/*
		递归
	 */
	//if root == nil {
	//	return 0
	//} else {
	//	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
	//}

	/*
		非递归
	 */
	if root == nil {
		return 0
	}
	result := 0
	queue1 := queue.New()
	queue1.Add(root)
	for queue1.Length() != 0 {
		result++
		queue2 := queue.New()
		for queue1.Length() != 0 {
			node := queue1.Remove().(*TreeNode)
			if node.Left != nil {
				queue2.Add(node.Left)
			}
			if node.Right != nil {
				queue2.Add(node.Right)
			}
		}
		queue1 = queue2
	}
	return result
}

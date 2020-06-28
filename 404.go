package main

import "fmt"

func main() {
	treeNode15 := TreeNode{
		Val:15,
	}
	treeNode7 := TreeNode{
		Val:7,
	}
	treeNode20 := TreeNode{
		Val:20,
		Left:&treeNode15,
		Right:&treeNode7,
	}
	treeNode9 := TreeNode{
		Val:9,
	}
	root := TreeNode{
		Val:3,
		Left:&treeNode9,
		Right:&treeNode20,
	}
	fmt.Println(sumOfLeftLeaves(&root))
}


//计算给定二叉树的所有左叶子之和。
//
//示例：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-left-leaves
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// sum-of-left-leaves
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	calu(root, &sum)
	return sum
}

func calu(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 判断是否是左叶子
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	calu(root.Left, sum)
	calu(root.Right, sum)
}

func sumOfLeftLeaves2(root *TreeNode) int {
	sum := 0
	var funcCalu func(root *TreeNode, sum *int)
	funcCalu = func(root *TreeNode, sum *int) {
		if root == nil {
			return
		}
		// 判断是否是左叶子
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			*sum += root.Left.Val
		}
		funcCalu(root.Left, sum)
		funcCalu(root.Right, sum)
	}

	funcCalu(root, &sum)
	return sum
}

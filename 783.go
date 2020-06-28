package main

import (
	"math"
	"fmt"
)


func main() {
	root := &TreeNode{
		Val:4,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:1,
			},
			Right:&TreeNode{
				Val:3,
			},
		},
		Right:&TreeNode{
			Val:6,
		},
	}
	fmt.Println(minDiffInBST(root))
}
//给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。
//
//示例：
//
//输入: root = [4,2,6,1,3,null,null]
//输出: 1
//解释:
//注意，root是树结点对象(TreeNode object)，而不是数组。
//
//给定的树 [4,2,6,1,3,null,null] 可表示为下图:
//
//          4
//        /   \
//      2      6
//     / \
//    1   3
//
//最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
func minDiffInBST(root *TreeNode) int {
	nums := inOrder(root)
	min := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		if tmp := nums[i+1] - nums[i]; tmp < min {
			min = tmp
		}
	}
	return min
}

func inOrder(root *TreeNode) []int {
	res := make([]int,0)
	if root != nil {
		if root.Left != nil {
			res = append(res, inOrder(root.Left)...)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			res = append(res, inOrder(root.Right)...)
		}
	}
	return res

}

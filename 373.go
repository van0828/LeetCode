package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{3}, 3))
}

//给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。
//
//定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。
//
//找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
//
//示例 1:
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//示例 2:
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//示例 3:
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//链接：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums

//type KSmallestPairsItem struct {
//	value    []int
//	priority int
//	index    int
//}
//
//type KSmallestPairsPQ []*KSmallestPairsItem
//
//func (pq KSmallestPairsPQ) Len() int { return len(pq) }
//
//func (pq KSmallestPairsPQ) Less(i, j int) bool {
//	return pq[i].priority < pq[j].priority
//}
//
//func (pq KSmallestPairsPQ) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *KSmallestPairsPQ) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*KSmallestPairsItem)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *KSmallestPairsPQ) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	item.index = -1
//	*pq = old[0 : n-1]
//	return item
//}
//
//func (pq *KSmallestPairsPQ) update(item *KSmallestPairsItem, value []int, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}
//
//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	pq := make(KSmallestPairsPQ, len(nums1)*len(nums2))
//	i := 0
//	for _, num1 := range nums1 {
//		for _, num2 := range nums2 {
//			pq[i] = &KSmallestPairsItem{
//				value:    []int{num1, num2},
//				priority: num1 + num2,
//				index:    i,
//			}
//			i++
//		}
//	}
//	heap.Init(&pq)
//	result := make([][]int, 0)
//	if k > len(nums1)*len(nums2) {
//		k = len(nums1)*len(nums2)
//	}
//	for k > 0 {
//		result = append(result, heap.Pop(&pq).(*KSmallestPairsItem).value)
//		k--
//	}
//	return result
//}

type KSmallestPairsItem struct {
	value []int
	sum   int
}

type KSmallestPairsHeap []*KSmallestPairsItem

func (h KSmallestPairsHeap) Len() int           { return len(h) }
func (h KSmallestPairsHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h KSmallestPairsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KSmallestPairsHeap) Push(x interface{}) {
	*h = append(*h, x.(*KSmallestPairsItem))
}

func (h *KSmallestPairsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := make(KSmallestPairsHeap, len(nums1)*len(nums2))
	i := 0
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			pq[i] = &KSmallestPairsItem{
				value: []int{num1, num2},
				sum:   num1 + num2,
			}
			i++
		}
	}
	heap.Init(&pq)
	result := make([][]int, 0)
	if k > len(nums1)*len(nums2) {
		k = len(nums1) * len(nums2)
	}
	for k > 0 {
		result = append(result, heap.Pop(&pq).(*KSmallestPairsItem).value)
		k--
	}
	return result
}

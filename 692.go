package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequentWords([]string{"the", "day", "is", "is", "day", "the", "the", "sunny", "day", "is"}, 3))
}

//给一非空的单词列表，返回前 k 个出现次数最多的单词。
//
//返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
//
//示例 1：
//
//输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
//
//示例 2：
//
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
//
//注意：
//
//假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
//输入的单词均由小写字母组成。
//
//链接：https://leetcode-cn.com/problems/top-k-frequent-words
//type Item struct {
//	value    string
//	priority int
//	index    int
//}
//
//type KSmallestPairsPQ []*Item
//
//func (pq KSmallestPairsPQ) Len() int { return len(pq) }
//
//func (pq KSmallestPairsPQ) Less(i, j int) bool {
//	if pq[i].priority == pq[j].priority {
//		return pq[i].value < pq[j].value
//	}
//	return pq[i].priority > pq[j].priority
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
//	item := x.(*Item)
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
//func (pq *KSmallestPairsPQ) update(item *Item, value string, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}
//
//func topKFrequentWords(words []string, k int) []string {
//	timesMap := make(map[string]int, 0)
//	for _, word := range words {
//		timesMap[word]++
//	}
//	pq := make(KSmallestPairsPQ, len(timesMap))
//	i := 0
//	for value, priority := range timesMap {
//		pq[i] = &Item{
//			value:    value,
//			priority: priority,
//			index:    i,
//		}
//		i++
//	}
//	heap.Init(&pq)
//	result := make([]string, 0)
//	for k > 0 {
//		result = append(result, heap.Pop(&pq).(*Item).value)
//		k--
//	}
//	return result
//}


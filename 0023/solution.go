package _023

import "container/heap"

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *hp) Pop() any          { elem := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return elem }

func _mergeKLists(lists []*ListNode) *ListNode {

	h := hp{}
	for _, p := range lists {
		if p != nil {
			heap.Push(&h, p)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for len(h) > 0 {
		next := heap.Pop(&h).(*ListNode)
		cur.Next = next
		cur = cur.Next

		if next.Next != nil {
			heap.Push(&h, next.Next)
		}
	}
	return dummy.Next
}

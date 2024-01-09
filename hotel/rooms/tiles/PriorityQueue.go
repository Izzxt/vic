package tiles

import (
	"container/heap"

	"github.com/Izzxt/vic/core"
)

type Item struct {
	value    core.IRoomTile // The value of the item; arbitrary.
	priority int            // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Contains(x any) bool {
	for _, item := range *pq {
		if item.value == x {
			return true
		}
	}
	return false
}

func (pq *PriorityQueue) Get(x int) *Item {
	for _, item := range *pq {
		if item.index == x {
			return item
		}
	}
	return nil
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Remove(x any) {
	for i, item := range *pq {
		if item.value == x {
			heap.Remove(pq, i)
		}
	}
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value core.IRoomTile, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

package queue

import (
	// "fmt"
	"math/rand"
	// "github.com/k0kubun/pp"
	"github.com/arccoza/go-queue/generic"
)

var ItemQueueTruncFreq int = 33
var itemQueueTruncCount = 0 // For testing only

type Item = generic.Type

type ItemQueue []Item

func NewItemQueue(cap int) ItemQueue {
	return make(ItemQueue, 0, cap)
}

func (q ItemQueue) Enqueue(items ...Item) ItemQueue {
	return append(q.trunc(2*(len(q)+len(items))), items...)
}

func (q ItemQueue) trunc(size int) ItemQueue {
	// Truncate queue
	if rand.Intn(ItemQueueTruncFreq) == ItemQueueTruncFreq/2 {
		q = append(make(ItemQueue, 0, size), q...)
		// pp.Println("TRUNC")
		itemQueueTruncCount++
	}
	return q
}

func (q ItemQueue) Dequeue(items []Item) (ItemQueue, int) {
	count := copy(items, q)
	return q[count:], count
}

func (q ItemQueue) First() Item {
	return q[0]
}

func (q ItemQueue) Last() Item {
	return q[len(q)-1]
}

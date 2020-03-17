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
	return append(q, items...)
}

func (q ItemQueue) Dequeue(items []Item) (ItemQueue, int) {
	count := copy(items, q)
	return q[count:].trunc(count), count
}

func (q ItemQueue) First() Item {
	return q[0]
}

func (q ItemQueue) Last() Item {
	return q[len(q)-1]
}

func (q ItemQueue) trunc(count int) ItemQueue {
	if count == 0 { return q }
	// Truncate queue
	if odds := (ItemQueueTruncFreq / count) + 2; rand.Intn(odds) == odds/2 {
		q = append(make(ItemQueue, 0, len(q) + count), q...)
		// pp.Println("TRUNC")
		itemQueueTruncCount++
	}
	return q
}

func minInt(a, b int) int {
	if a < b { return a }
	return b
}

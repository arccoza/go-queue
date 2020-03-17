package queue

import (
	// "fmt"
	"math/rand"
	// "github.com/k0kubun/pp"
	"github.com/arccoza/go-queue/generic"
)

var ItemQueueTruncFreq int = 33

type Item = generic.Type

type ItemQueue []Item

func NewItemQueue(cap int) ItemQueue {
	return make(ItemQueue, 0, cap)
}

func (q ItemQueue) Enqueue(items ...Item) ItemQueue {
	// Truncate queue
	if rand.Intn(ItemQueueTruncFreq) == ItemQueueTruncFreq/2 {
		q = append(make(ItemQueue, 0, 2*(len(q)+len(items))), q...)
	}
	return append(q, items...)
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

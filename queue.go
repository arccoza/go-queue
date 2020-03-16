package queue

import (
	// "fmt"
	"math/rand"
	// "github.com/k0kubun/pp"
)

var TruncFreq int = 33

type Queue []interface{}

func New(cap int) Queuer {
	return make(Queue, 0, cap)
}

func (q Queue) Enqueue(items ...interface{}) Queue {
	// Truncate queue
	if rand.Intn(TruncFreq) == TruncFreq/2 {
		q = append(make(Queue, 0, 2*(len(q)+len(items))), q...)
	}
	return append(q, items...)
}

func (q Queue) Dequeue(items []interface{}) (Queue, int) {
	count := copy(items, q)
	return q[count:], count
}

func (q Queue) First() interface{} {
	return q[0]
}

func (q Queue) Last() interface{} {
	return q[len(q)-1]
}

type Queuer interface {
	Enqueue(...interface{}) Queuer
	Dequeue([]interface{}) (Queuer, int)
	First() interface{}
	Last() interface{}
}

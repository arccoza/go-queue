package queue

import (
	"testing"
	// "github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	q := NewItemQueue(16)
	want := ItemQueue(src[:6])
	got := q.Enqueue(src[:6]...)
	t.Log(want)
	t.Log(got)
	assert.Equalf(t, want, got, "Queue should have %v elements", want)
}

func TestDequeue(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	gotDest := make([]interface{}, 4)
	q := NewItemQueue(16)
	q = q.Enqueue(src[:6]...)
	wantQ, wantCount, wantDest := ItemQueue(src[4:6]), 4, src[:4]
	gotQ, gotCount := q.Dequeue(gotDest)
	t.Log(wantQ, wantCount, wantDest)
	t.Log(gotQ, gotCount, gotDest)
	assert.Equalf(t, wantQ, gotQ, "Queue should equal %v", wantQ)
	assert.Equalf(t, wantCount, gotCount, "Count should equal %v", wantCount)
	assert.Equalf(t, wantDest, gotDest, "Dequeued items should equal %v", wantDest)
}

func TestFirst(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	q := NewItemQueue(16)
	want := ItemQueue(src[:6])[0]
	got := q.Enqueue(src[:6]...).First()
	t.Log(want)
	t.Log(got)
	assert.Equalf(t, want, got, "First should be %v", want)
}

func TestLast(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	q := NewItemQueue(16)
	want := ItemQueue(src[:6])[5]
	got := q.Enqueue(src[:6]...).Last()
	t.Log(want)
	t.Log(got)
	assert.Equalf(t, want, got, "Last should be %v", want)
}

func TestTrunc(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	dest := make([]interface{}, 60)
	q := NewItemQueue(16).Enqueue(src...).Enqueue(src...).Enqueue(src...).Enqueue(src...)

	q2, _ := q.Dequeue(dest)
	for i := 1; i < 24; i++ {
		q2 = q2.Enqueue(i)
	}

	t.Log(q, len(q), cap(q))
	t.Log(q2, len(q2), cap(q2))
	assert.Lessf(t, cap(q2), cap(q) + 24, "Queue should be truncated %v %v", q, q2)
	t.Log(itemQueueTruncCount)
	assert.Greaterf(t, itemQueueTruncCount, 0, "Truncate must have run at least once: %v", itemQueueTruncCount)
}

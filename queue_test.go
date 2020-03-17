package queue

import (
	"fmt"
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
	assert.Equal(t, want, got, fmt.Sprintf("Queue should have %v elements", want))
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
	assert.Equal(t, wantQ, gotQ, fmt.Sprintf("Queue should equal %v", wantQ))
	assert.Equal(t, wantCount, gotCount, fmt.Sprintf("Count should equal %v", wantCount))
	assert.Equal(t, wantDest, gotDest, fmt.Sprintf("Dequeued items should equal %v", wantDest))
}

func TestFirst(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	q := NewItemQueue(16)
	want := ItemQueue(src[:6])[0]
	got := q.Enqueue(src[:6]...).First()
	t.Log(want)
	t.Log(got)
	assert.Equal(t, want, got, fmt.Sprintf("First should be %v", want))
}

func TestLast(t *testing.T) {
	src := []interface{}{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	q := NewItemQueue(16)
	want := ItemQueue(src[:6])[5]
	got := q.Enqueue(src[:6]...).Last()
	t.Log(want)
	t.Log(got)
	assert.Equal(t, want, got, fmt.Sprintf("Last should be %v", want))
}

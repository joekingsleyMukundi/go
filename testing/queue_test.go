// use `go test -v ./demo/testing` to run tests
package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("failed to append item %v to queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("should not be able to add any more items to the array")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get an item from the queue")
		}
		if item != i {
			t.Errorf("got items i the wrong order: %v want %v", item, i)
		}
	}
	//queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("shounld not be more items in the queue. got: %v", item)
	}
}

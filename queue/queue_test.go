package queue

import "testing"

func TestNew(t *testing.T) {
	q := New()
	if q == nil {
		t.Errorf("Expected queue, got nil")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("queue: Expected size of 3, got %d", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	element, err := q.Dequeue()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if element != 1 {
		t.Errorf("element: Expected 1, got %d", element)
	}
	if q.Size() != 2 {
		t.Errorf("queue: Expected size of 2, got %d", q.Size())
	}
}

func TestQueue_Pop(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	element, err := q.Pop()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if element != 3 {
		t.Errorf("element: Expected 3, got %d", element)
	}
	if q.Size() != 2 {
		t.Errorf("queue: Expected size of 2, got %d", q.Size())
	}
}

func TestQueue_PopEmpty(t *testing.T) {
	q := New()
	_, err := q.Pop()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestQueue_Push(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Push(3)
	want := 3
	got, _ := q.Peek()
	if got != want {
		t.Errorf("element: Expected %d, got %d", want, got)
	}
}

func TestQueue_PushEmpty(t *testing.T) {
	q := New()
	q.Push(1)
	want := 1
	got, _ := q.Peek()
	if got != want {
		t.Errorf("element: Expected %d, got %d", want, got)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	element, err := q.Peek()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if element != 1 {
		t.Errorf("element: Expected 1, got %d", element)
	}
	if q.Size() != 3 {
		t.Errorf("queue: Expected size of 3, got %d", q.Size())
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Errorf("queue: Expected empty queue")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("queue: Expected non-empty queue")
	}
}

func TestQueue_Size(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("queue: Expected size of 3, got %d", q.Size())
	}
}

func TestQueue_Clear(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Clear()
	if q.Size() != 0 {
		t.Errorf("queue: Expected size of 0, got %d", q.Size())
	}
}

func TestQueue_Values(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	values := q.Values()
	if len(values) != 3 {
		t.Errorf("values: Expected length of 3, got %d", len(values))
	}
}

func TestQueue_Contains(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if !q.Contains(2) {
		t.Errorf("contains: Expected true")
	}
	if q.Contains(4) {
		t.Errorf("contains: Expected false")
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	q := New()
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestQueue_PeekEmpty(t *testing.T) {
	q := New()
	_, err := q.Peek()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestQueue_ContainsEmpty(t *testing.T) {
	q := New()
	if q.Contains(1) {
		t.Errorf("Expected false")
	}
}

func TestQueue_ValuesEmpty(t *testing.T) {
	q := New()
	values := q.Values()
	if len(values) != 0 {
		t.Errorf("values: Expected length of 0, got %d", len(values))
	}
}

func TestQueue_Remove(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Remove(2)
	if q.Contains(2) {
		t.Errorf("contains: Expected false")
	}
	if q.Size() != 2 {
		t.Errorf("queue: Expected size of 2, got %d", q.Size())
	}
}

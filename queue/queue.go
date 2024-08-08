package queue

import "fmt"

type Queue struct {
	elements []interface{}
}

// New creates a new queue
func New() *Queue {
	return &Queue{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(element interface{}) {
	q.elements = append(q.elements, element)
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

// Peek returns the element at the front of the queue
func (q *Queue) Peek() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.elements)
}

// Clear removes all elements from the queue
func (q *Queue) Clear() {
	q.elements = []interface{}{}
}

// Values returns all elements in the queue
func (q *Queue) Values() []interface{} {
	return q.elements
}

// Contains returns true if the queue contains the given element
func (q *Queue) Contains(element interface{}) bool {
	for _, e := range q.elements {
		if e == element {
			return true
		}
	}
	return false
}

// Remove removes elements from a queue
func (q *Queue) Remove(elements ...interface{}) {
	for _, element := range elements {
		for i, v := range q.elements {
			if v == element {
				q.elements = append(q.elements[:i], q.elements[i+1:]...)
			}
		}
	}
}

// Pop removes and returns the last element from the queue
func (q *Queue) Pop() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	element := q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
	return element, nil
}

// Push adds an element to the front of the queue
func (q *Queue) Push(element interface{}) {
	q.elements = append([]interface{}{element}, q.elements...)
}

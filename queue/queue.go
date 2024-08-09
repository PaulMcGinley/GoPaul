// Created by Paul F. McGinley Aug 8, 2024

// Package queue provides a object queue implementation.
// Unlike a fixed-size array implementation, this queue implementation is dynamic and can grow as needed.
// The Queue type can be used as a fifo queue, lifo queue, or priority queue.
package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	elements []interface{}
	lock     sync.Mutex
}

// New creates a new queue
func New() *Queue {
	return &Queue{}
}

// Length returns the number of elements in the queue
func (q *Queue) Length() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.elements)
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.elements) == 0
}

// Clear removes all elements from the queue
func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = []interface{}{}
}

// Values returns all elements in the queue
func (q *Queue) Values() []interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.elements
}

// Contains returns true if the queue contains the given element
func (q *Queue) Contains(element interface{}) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, e := range q.elements {
		if e == element {
			return true
		}
	}
	return false
}

// Remove removes elements from a queue
func (q *Queue) Remove(elements ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, element := range elements {
		for i, v := range q.elements {
			if v == element {
				q.elements = append(q.elements[:i], q.elements[i+1:]...)
			}
		}
	}
}

// RemoveMultiple removes multiple elements from a queue
func (q *Queue) RemoveMultiple(elements ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, element := range elements {
		for i := 0; i < len(q.elements); i++ {
			if q.elements[i] == element {
				q.elements = append(q.elements[:i], q.elements[i+1:]...)
				i--
			}
		}
	}
}

// Swap swaps two elements in the queue
func (q *Queue) Swap(a, b int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if a < 0 || a >= len(q.elements) {
		return fmt.Errorf("index i (%d) out of range", a)
	}
	if b < 0 || b >= len(q.elements) {
		return fmt.Errorf("index j (%d) out of range", b)
	}

	q.elements[a], q.elements[b] = q.elements[b], q.elements[a]
	return nil
}

// Reverse reverses the order of elements in the queue
func (q *Queue) Reverse() {
	q.lock.Lock()
	defer q.lock.Unlock()

	for a, b := 0, len(q.elements)-1; a < b; a, b = a+1, b-1 {
		q.elements[a], q.elements[b] = q.elements[b], q.elements[a]
	}
}

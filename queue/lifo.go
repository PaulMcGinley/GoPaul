package queue

import "fmt"

// DequeueLast removes and returns the last element from the queue
func (q *Queue) DequeueLast() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	element := q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
	return element, nil
}

// DequeueLastMultiple removes and returns the last n elements from the queue
func (q *Queue) DequeueLastMultiple(n int) ([]interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if n > len(q.elements) {
		return nil, fmt.Errorf("queue does not have %d elements", n)
	}
	elements := q.elements[len(q.elements)-n:]
	q.elements = q.elements[:len(q.elements)-n]
	return elements, nil
}

// EnqueueFirst adds an element to the front of the queue
func (q *Queue) EnqueueFirst(element interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = append([]interface{}{element}, q.elements...)
}

// EnqueueFirstMultiple adds multiple elements to the front of the queue
func (q *Queue) EnqueueFirstMultiple(elements ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = append(elements, q.elements...)
}

package queue

import "fmt"

func (q *Queue) Enqueue(element interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = append(q.elements, element)
}

// EnqueueMultiple adds multiple elements to the end of the queue
func (q *Queue) EnqueueMultiple(elements ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = append(q.elements, elements...)
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.IsEmpty() {
		return nil, IsEmptyError
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

// DequeueMultiple removes n elements from the front of the queue
func (q *Queue) DequeueMultiple(n int) ([]interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if n > q.Length() {
		return nil, fmt.Errorf("queue does not have %d elements", n)
	}

	elements := q.elements[:n]
	q.elements = q.elements[n:]
	return elements, nil
}

// DequeueAll removes all elements from the queue
func (q *Queue) DequeueAll() ([]interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	elements := q.elements
	q.elements = []interface{}{} // Empty the queue

	if len(elements) == 0 {
		return nil, IsEmptyError
	}
	return elements, nil
}

// DequeueN removes n elements from the front of the queue
func (q *Queue) DequeueN(n int) ([]interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if n > len(q.elements) {
		return nil, fmt.Errorf("queue does not have %d elements", n)
	}
	elements := q.elements[:n]
	q.elements = q.elements[n:]
	return elements, nil
}

// DequeueWhile removes elements from the front of the queue while the predicate is true
func (q *Queue) DequeueWhile(predicate func(interface{}) bool) []interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	var elements []interface{}
	for i, element := range q.elements {
		if predicate(element) {
			elements = append(elements, element)
		} else {
			q.elements = q.elements[i:]
			break
		}
	}
	return elements
}

// Peek returns the element at the front of the queue
func (q *Queue) Peek() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.elements) == 0 {
		return nil, IsEmptyError
	}
	return q.elements[0], nil
}

// PeekMultiple returns the first n elements from the queue
func (q *Queue) PeekMultiple(n int) ([]interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if n > len(q.elements) {
		return nil, fmt.Errorf("queue does not have %d elements", n)
	}
	return q.elements[:n], nil
}

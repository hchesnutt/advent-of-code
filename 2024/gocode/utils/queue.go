package utils

import (
	"container/list"
)

// Queue represents a queue data structure using a linked list
type Queue struct {
	list *list.List
}

// NewQueue creates and initializes a new Queue
func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.list.PushBack(value)
}

// Dequeue removes and returns the element from the front of the queue
// Returns nil if the queue is empty
func (q *Queue) Dequeue() interface{} {
	front := q.list.Front()
	if front != nil {
		q.list.Remove(front)
		return front.Value
	}
	return nil
}

// Peek returns the element at the front of the queue without removing it
// Returns nil if the queue is empty
func (q *Queue) Peek() interface{} {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.list.Len()
}

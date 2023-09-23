package lineards

import (
	"fmt"
	"strings"
	"sync"
)

// Queue is a linear data structure that follows the FIFO (First In First Out) principle. It is goroutine safe.
type Queue[T any] struct {
	queue []T
	mutex *sync.RWMutex
	cond  *sync.Cond
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{
		queue: []T{},
		mutex: &sync.RWMutex{},
	}
	q.cond = sync.NewCond(q.mutex)
	return q
}

// Enqueue adds an item to the queue.
func (q *Queue[T]) Enqueue(item ...T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item...)
	q.cond.Signal()
}

// Dequeue removes an item from the queue.
func (q *Queue[T]) Dequeue() T {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for len(q.queue) == 0 {
		q.cond.Wait()
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	return item
}

// DequeueAll removes all items from the queue.
func (q *Queue[T]) DequeueAll() []T {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	items := q.queue
	q.queue = []T{}
	return items
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.queue) == 0
}

// Peek returns the first item in the queue without removing it.
func (q *Queue[T]) Peek() T {
	for len(q.queue) == 0 {
		q.cond.Wait()
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.queue[0]
}

// PeekAll returns all items in the queue without removing them.
func (q *Queue[T]) PeekAll() []T {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.queue
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.queue)
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	var str []string
	for _, item := range q.queue {
		str = append(str, fmt.Sprintf("%v", item))
	}

	return strings.Join(str, ", ")
}

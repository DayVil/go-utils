package lineards

import (
	"fmt"
	"strings"
	"sync"
)

type Queue[T any] struct {
	queue []T
	mutex *sync.RWMutex
	cond  *sync.Cond
}

func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{
		queue: []T{},
		mutex: &sync.RWMutex{},
	}
	q.cond = sync.NewCond(q.mutex)
	return q
}

func (q *Queue[T]) Enqueue(item ...T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item...)
	q.cond.Signal()
}

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

func (q *Queue[T]) DequeueAll() []T {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	items := q.queue
	q.queue = []T{}
	return items
}

func (q *Queue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.queue) == 0
}

func (q *Queue[T]) Peek() T {
	for len(q.queue) == 0 {
		q.cond.Wait()
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.queue[0]
}

func (q *Queue[T]) PeekAll() []T {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.queue
}

func (q *Queue[T]) Len() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.queue)
}

func (q *Queue[T]) String() string {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	str := []string{}
	for _, item := range q.queue {
		str = append(str, fmt.Sprintf("%v", item))
	}

	return strings.Join(str, ", ")
}

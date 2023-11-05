package lineards

import _ "sync"

type DoubleSidedQueue[T any] struct {
	Queue[T]
}

func NewDoubleSidedQueue[T any]() *DoubleSidedQueue[T] {
	q := NewQueue[T]()
	return &DoubleSidedQueue[T]{*q}
}

func (dq *DoubleSidedQueue[T]) DequeueBack() T {
	dq.Lock()
	defer dq.Unlock()

	for len(dq.queue) == 0 {
		dq.Wait()
	}

	lastItem := len(dq.queue) - 1
	item := dq.queue[lastItem]
	dq.queue = dq.queue[:lastItem]

	return item
}

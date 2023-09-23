package lineards

import (
	"fmt"
	"strings"
	"sync"
)

type Stack[T any] struct {
	stack []T
	mutex *sync.RWMutex
	cond  *sync.Cond
}

func NewStack[T any]() *Stack[T] {
	s := &Stack[T]{
		stack: []T{},
		mutex: &sync.RWMutex{},
	}
	s.cond = sync.NewCond(s.mutex)
	return s
}

func (s *Stack[T]) Enqueue(item ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, i := range item {
		s.stack = append([]T{i}, s.stack...)
		s.cond.Signal()
	}
}

func (s *Stack[T]) Dequeue() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for len(s.stack) == 0 {
		s.cond.Wait()
	}

	item := s.stack[0]
	s.stack = s.stack[1:]

	return item
}

func (s *Stack[T]) DequeueAll() []T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	items := s.stack
	s.stack = []T{}
	return items
}

func (s *Stack[T]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.stack) == 0
}

func (s *Stack[T]) Peek() T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for len(s.stack) == 0 {
		s.cond.Wait()
	}
	return s.stack[0]
}

func (s *Stack[T]) PeekAll() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.stack
}

func (s *Stack[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.stack)
}

func (s *Stack[T]) String() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var str []string
	for _, item := range s.stack {
		str = append(str, fmt.Sprintf("%v", item))
	}

	return strings.Join(str, ", ")
}

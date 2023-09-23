package lineards

import (
	"fmt"
	"strings"
	"sync"
)

// Stack is a linear data structure that follows the LIFO (Last In First Out) principle. It is goroutine safe.
type Stack[T any] struct {
	stack []T
	mutex *sync.RWMutex
	cond  *sync.Cond
}

// NewStack creates a new stack.
func NewStack[T any]() *Stack[T] {
	s := &Stack[T]{
		stack: []T{},
		mutex: &sync.RWMutex{},
	}
	s.cond = sync.NewCond(s.mutex)
	return s
}

// Enqueue Push adds an item to the stack.
func (s *Stack[T]) Enqueue(item ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, i := range item {
		s.stack = append([]T{i}, s.stack...)
		s.cond.Signal()
	}
}

// Dequeue Pop removes an item from the stack.
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

// DequeueAll removes all items from the stack.
func (s *Stack[T]) DequeueAll() []T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	items := s.stack
	s.stack = []T{}
	return items
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.stack) == 0
}

// Peek returns the top item from the stack without removing it.
func (s *Stack[T]) Peek() T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for len(s.stack) == 0 {
		s.cond.Wait()
	}
	return s.stack[0]
}

// PeekAll returns all items from the stack without removing them.
func (s *Stack[T]) PeekAll() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.stack
}

// Len returns the number of items in the stack.
func (s *Stack[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.stack)
}

// String returns a string representation of the stack.
func (s *Stack[T]) String() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var str []string
	for _, item := range s.stack {
		str = append(str, fmt.Sprintf("%v", item))
	}

	return strings.Join(str, ", ")
}

package lineards

// Stack is a linear data structure that follows the LIFO (Last In First Out) principle. It is goroutine safe.
type Stack[T any] struct {
	Queue[T]
}

// NewStack creates a new stack.
func NewStack[T any]() *Stack[T] {
	s := NewQueue[T]()
	return &Stack[T]{*s}
}

// Enqueue Push adds an item to the stack.
func (s *Stack[T]) Enqueue(item ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, i := range item {
		s.queue = append([]T{i}, s.queue...)
		s.cond.Signal()
	}
}

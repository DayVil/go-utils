package monads

import (
	"errors"
	"sync"
)

// TODO: add tests

// Optional is a monad that represents a value that may or may not exist. Optional is go-routine safe.
type Optional[T any] struct {
	value *T
	mutex *sync.RWMutex
}

// NewOptional creates a new Optional
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		&value,
		&sync.RWMutex{},
	}
}

// NewEmptyOptional creates a new empty Optional
func NewEmptyOptional[T any]() Optional[T] {
	return Optional[T]{nil, &sync.RWMutex{}}
}

// SetValue sets the value of the Optional
func (o Optional[T]) SetValue(value T) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	*o.value = value
}

// IsPresent checks if the Optional has a value
func (o Optional[T]) IsPresent() bool {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return o.value != nil
}

// Get returns the value of the Optional
func (o Optional[T]) Get() (T, error) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	if o.value == nil {
		var zero T
		return zero, errors.New("no value present")
	}

	val := *o.value
	return val, nil
}

// MustGet returns the value of the Optional, panicking if there is no value
func (o Optional[T]) MustGet() T {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	val, err := o.Get()
	if err != nil {
		panic(err)
	}

	return val
}

// IfPresent executes the given function if the Optional has a value
func (o Optional[T]) IfPresent(f func(T)) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	if o.value != nil {
		val := *o.value
		f(val)
	}
}

// OrElse returns the value of the Optional if it exists, otherwise it returns the given value
func (o Optional[T]) OrElse(value T) T {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	if o.value == nil {
		return value
	}

	val := *o.value
	return val
}

// OrElseGet returns the value of the Optional if it exists, otherwise it returns the value returned by the given function
func (o Optional[T]) OrElseGet(f func() T) T {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	if o.value == nil {
		return f()
	}

	val := *o.value
	return val
}

package lineards

// LinearDS is an interface for linear data structures such as stacks and queues. Everything is goroutine safe.
type LinearDS[T any] interface {
	// Enqueue Push adds an item to the data structure
	Enqueue(item ...T)

	// Dequeue Pop removes an item from the data structure
	Dequeue() T

	// DequeueAll Pops all items from the data structure
	DequeueAll() []T

	// IsEmpty checks if the data structure is empty
	IsEmpty() bool

	// Peek returns the top item from the data structure
	Peek() T

	// PeekAll returns all items from the data structure
	PeekAll() []T
}

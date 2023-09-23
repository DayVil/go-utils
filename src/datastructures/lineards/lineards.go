package lineards

type LinearDS[T any] interface {
	Enqueue(item ...T)
	Dequeue() T
	DequeueAll() []T
	IsEmpty() bool
	Peek() T
	PeekAll() []T
}

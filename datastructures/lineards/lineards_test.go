package lineards

import (
	"slices"
	"testing"
	"time"
)

func asyncEnqueue(fun func(...string), t time.Duration, ele ...string) {
	time.Sleep(t)
	fun(ele...)
}

func TestLinearDS(t *testing.T) {
	hello := "Hello"
	bye := "Bye"
	see := "See"
	fall := "Fall"
	elements := []string{hello, bye, see, fall}
	elementsR := slices.Clone(elements)
	slices.Reverse(elementsR)

	queue := NewQueue[string]()
	stack := NewStack[string]()
	doubleSidedQueue := NewDoubleSidedQueue[string]()
	fDoubleSidedQueue := NewDoubleSidedQueue[string]()

	go asyncEnqueue(queue.Enqueue, time.Second*3, elements...)
	go asyncEnqueue(stack.Enqueue, time.Second, elements...)
	go asyncEnqueue(doubleSidedQueue.Enqueue, time.Second*5, elements...)
	go asyncEnqueue(fDoubleSidedQueue.Enqueue, time.Second*2, elements...)

	for i := range elements {
		sEle := stack.Dequeue()
		qEle := queue.Dequeue()
		dqEle := doubleSidedQueue.Dequeue()
		fdqEle := fDoubleSidedQueue.DequeueBack()

		actualQVal := elements[i]
		if qEle != actualQVal {
			t.Fatalf("Queue: got %s but expected %s\n", qEle, actualQVal)
		}
		if qEle != dqEle {
			t.Fatalf("dqQueue: got %s but expected %s\n", dqEle, qEle)
		}

		actualSVal := elements[len(elements)-(i+1)]
		if sEle != actualSVal {
			t.Fatalf("Queue: got %s but expected %s\n", sEle, actualSVal)
		}
		if sEle != fdqEle {
			t.Fatalf("fdqQueue: got %s but expected %s\n", fdqEle, sEle)
		}
	}
}

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

	go asyncEnqueue(queue.Enqueue, time.Second*3, elements...)
	go asyncEnqueue(stack.Enqueue, time.Second, elements...)

	for i := range elements {
		sEle := stack.Dequeue()
		qEle := queue.Dequeue()

		actualQVal := elements[i]
		if qEle != actualQVal {
			t.Fatalf("Queue: got %s but expected %s\n", qEle, actualQVal)
		}

		actualSVal := elements[len(elements)-(i+1)]
		if sEle != actualSVal {
			t.Fatalf("Queue: got %s but expected %s\n", sEle, actualSVal)
		}
	}
}

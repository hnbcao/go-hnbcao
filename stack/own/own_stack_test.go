package own

import "testing"

var ownStack = NewOwnStack()

func BenchmarkStack_Push(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		ownStack.Push(i)
	}

}

func BenchmarkStack_Pop(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_, _ = ownStack.Pop()
	}
}

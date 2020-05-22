package slice

import (
	"testing"
)

var sliceStack = NewSliceStack()

func BenchmarkSliceStack_Push(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		sliceStack.Push(i)
	}

}

func BenchmarkSliceStack_Pop(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_, _ = sliceStack.Pop()
	}
}

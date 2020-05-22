package list

import (
	"testing"
)

var listStack = NewListStack()

func BenchmarkListStack_Push(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		listStack.Push(i)
	}

}

func BenchmarkListStack_Pop(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		_, _ = listStack.Pop()
	}
}

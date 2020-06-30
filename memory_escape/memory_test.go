package memory_escape

import (
	"fmt"
	"testing"
)

func TestEscapeMemoryObj1(t *testing.T) {
	obj1 := &MemoryObj1{}
	obj2 := &MemoryObj1{}
	fmt.Printf("%p\n", obj1)
	fmt.Printf("%p\n", obj2)
	t.Log(obj1 == obj2)
	fmt.Printf("%p\n", obj1)
	fmt.Printf("%p\n", obj2)
}

func TestEscapeMemoryObj2(t *testing.T) {
	obj1 := &MemoryObj2{}
	obj2 := &MemoryObj2{}
	fmt.Printf("%p\n", obj1)
	fmt.Printf("%p\n", obj2)
	t.Log(obj1 == obj2)
	fmt.Printf("%p\n", obj1)
	fmt.Printf("%p\n", obj2)
}

package slice

import (
	"fmt"
	"go-hnbcao/stack"
	"sync"
)

type Item interface{}

type Stack struct {
	items []Item
	lock  sync.RWMutex
}

func (s *Stack) Print() {
	fmt.Println(s)
}

func (s *Stack) Push(item Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (item Item, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) < 1 {
		err = stack.NoItemError
	} else {
		item = s.items[len(s.items)-1]
		s.items = s.items[0 : len(s.items)-1]
	}
	return
}

func NewSliceStack() *Stack {
	s := &Stack{}
	s.items = []Item{}
	return s
}

package list

import (
	"container/list"
	"fmt"
	"go-hnbcao/stack"
	"sync"
)

type Stack struct {
	items *list.List
	lock  sync.RWMutex
}

func (s *Stack) Print() {
	fmt.Println(s)
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items.PushBack(item)
}

func (s *Stack) Pop() (item interface{}, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items.Len() < 1 {
		err = stack.NoItemError
	} else {
		tmp := s.items.Front()
		s.items.Remove(tmp)
		item = tmp.Value
	}
	return
}

func NewListStack() *Stack {
	s := &Stack{}
	s.items = list.New()
	return s
}

package own

import (
	"fmt"
	"go-hnbcao/stack"
	"sync"
)

type node struct {
	prev  *node
	value interface{}
}

type Stack struct {
	items *node
	lock  sync.RWMutex
}

func (s *Stack) Print() {
	fmt.Println(s)
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	nd := &node{value: item}
	nd.prev = s.items
	s.items = nd
}

func (s *Stack) Pop() (item interface{}, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.items == nil {
		err = stack.NoItemError
	} else {
		tmp := s.items
		s.items = tmp.prev
		item = tmp.value
	}
	return
}

func NewOwnStack() *Stack {
	s := &Stack{}
	return s
}

package gocore

import (
	"container/list"
	"sync"
)

type Stack struct {
	data *list.List
}

func NewStack() *Stack {
	return &Stack{data: list.New()}
}

func (s *Stack) Top() interface{} {
	if e := s.data.Front(); e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	s.data.PushBack(data)
}

func (s *Stack) Pop() interface{} {
	if e := s.data.Front(); e != nil {
		s.data.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.data.Len()
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Swap(other *Stack) {
	tmp := other.data
	other.data = s.data
	s.data = tmp
}

/********* safe **********/

type SafeStack struct {
	mu    *sync.RWMutex
	stack *Stack
}

func NewSafeStack() *SafeStack {
	return &SafeStack{
		mu:    new(sync.RWMutex),
		stack: NewStack(),
	}
}

func (ss *SafeStack) Top() interface{} {
	ss.mu.RLock()
	defer ss.mu.RLock()

	return ss.stack.Top()
}

func (ss *SafeStack) Push(data interface{}) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.stack.Push(data)
}

func (ss *SafeStack) Pop() interface{} {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	return ss.Pop()
}

func (ss *SafeStack) Len() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()

	return ss.stack.Len()
}

func (ss *SafeStack) Empty() bool {
	ss.mu.RLock()
	defer ss.mu.RUnlock()

	return ss.stack.Empty()
}

func (ss *SafeStack) Swap(other *SafeStack) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()

	ss.stack.Swap(other.stack)
}

package gocore

import (
	cl "container/list"
)

/**	
 * List
 * 
 * Back() *cl.Element
 * Clear()
 * Front() *cl.Element
 * Init() *cl.List
 * InsertAfter(v interface{}, mark *cl.Element) *cl.Element
 * InsertBefore(v interface{}, mark *cl.Element) *cl.Element
 * Len() int
 * MoveAfter(e *cl.Element, mark *cl.Element)
 * MoveBefore(e *cl.Element, mark *cl.Element)
 * MoveToBack(e *cl.Element)
 * MoveToFront(e *cl.Element)
 * PopBack()
 * PopFront()
 * PushBack(v interface{}) *cl.Element
 * PushBackList(other *cl.List)
 * PushFront(v interface{}) *cl.Element
 * PushFrontList(other *cl.List)
 * Remove(e *cl.Element) interface{}
 * RemoveEqual(v interface{})
**/

type List struct {
	cl.List
}

func NewListWithInit() *List {
	back := &List{}
	back.Init()
	return back
}

func NewList() *List {
	return &List{}
}

func (l *List) PopFront() {
	if l.Len() == 0 {
		return
	}
	l.Remove(l.Front())
}

func (l *List) PopBack() {
	if l.Len() == 0 {
		return
	}
	l.Remove(l.Back())
}

func (l *List) Clear() {
	for curr := l.Front(); curr != nil; curr = curr.Next() {
		l.Remove(curr)
	}
}

func (l *List) RemoveEqual(v interface{}) {
	for curr := l.Front(); curr != nil; {
		if curr.Value == v {
			tmp := curr
			curr = curr.Next()
			l.Remove(tmp)
			continue
		}
		curr = curr.Next()
	}
}

func (l *List) Unique() {
	first := l.Front()
	last := l.Back()
	if first == last || first == nil || last == nil {
		return
	}
	for itr := first; itr != last; {
		if itr.Value == last.Value {
			tmp := itr
			itr = itr.Next()
			l.Remove(tmp)
			continue
		} else {
			first = itr
		}
		itr = itr.Next()
	}
}

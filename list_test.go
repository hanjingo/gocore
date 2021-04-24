package gocore

import (
	cl "container/list"
	"testing"
)

func TestNewList(t *testing.T) {
	lold := cl.New()
	lold.Init()
	lnew := NewList()

	// insert
	lnew.Insert()
}

func TestInsert(t *testing.T) {
	lold := cl.New()
	lold.Init()
	e1 := lold.PushBack(1)
	e2 := lold.PushBack(2)

	lnew := NewList()
	lnew.PushBack(1)
	lnew.PushBack(2)

	lold.InsertAfter(3, e1)
	lnew.InsertAfter(1, 3)
}

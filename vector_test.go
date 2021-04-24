package gocore

import (
	"fmt"
	"testing"
)

// go test -v vector.go vector_test.go -test.run TestNewVector
func TestNewVector(t *testing.T) {
	v := NewVector()
	fmt.Printf("New vector with .data=%v, .finish=%d\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestEmpty
func TestEmpty(t *testing.T) {
	v := NewVector()
	fmt.Printf("first Empty():%v\n", v.Empty())
	v.PushBack("a")
	fmt.Printf("after PushBack(\"a\"), Empty():%v\n", v.Empty())
}

// go test -v vector.go vector_test.go -test.run TestAt
func TestAt(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	v.PushBack("b")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	fmt.Printf("v.At(0):%v\n", v.At(0))
	fmt.Printf("v.At(1):%v\n", v.At(1))
	fmt.Printf("v.At(2):%v\n", v.At(2))
}

// go test -v vector.go vector_test.go -test.run TestFront
func TestFront(t *testing.T) {
	v := NewVector()
	fmt.Printf("v.data:%v, v.finish:%v, v.Front():%v\n", v.data, v.finish, v.Front())
	v.PushBack("a")
	fmt.Printf("v.data:%v, v.finish:%v, v.Front():%v\n", v.data, v.finish, v.Front())
}

// go test -v vector.go vector_test.go -test.run TestBack
func TestBack(t *testing.T) {
	v := NewVector()
	fmt.Printf("v.data:%v, v.finish:%v, v.Back():%v\n", v.data, v.finish, v.Back())
	v.PushBack("a")
	fmt.Printf("v.data:%v, v.finish:%v, v.Back():%v\n", v.data, v.finish, v.Back())
}

// go test -v vector.go vector_test.go -test.run TestRange
func TestRange(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	v.PushBack("b")
	v.PushBack("c")
	v.PushBack("d")
	v.PushBack("e")
	v.PushBack("f")
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(0, 5):%v\n", v.data, v.finish, v.Range(0, 5))
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(5, 5):%v\n", v.data, v.finish, v.Range(5, 5))
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(5, 5):%v\n", v.data, v.finish, v.Range(5, 6))
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(0, 6):%v\n", v.data, v.finish, v.Range(0, 6))
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(4, 3):%v\n", v.data, v.finish, v.Range(4, 3))
	fmt.Printf("v.data:%v, v.finish:%v, v.Range(3, 5):%v\n", v.data, v.finish, v.Range(3, 5))
}


// go test -v vector.go vector_test.go -test.run TestPopBack
func TestPopBack(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	v.PopBack()
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestPopFront
func TestPopFront(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	v.PopFront()
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestPushBack
func TestPushBack(t *testing.T) {
	v := NewVector()
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	v.PushBack("a")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestPushFront
func TestPushFront(t *testing.T) {
	v := NewVector()
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	v.PushFront("a")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestErase
func TestErase(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	v.PushBack("b")
	v.PushBack("c")
	v.PushBack("d")
	v.PushBack("e")
	v.PushBack("f")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	n := v.Erase(0, 1)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(0, 1)\n", v.data, v.finish, n)
	n = v.Erase(0, 0)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(0, 0)\n", v.data, v.finish, n)
	n = v.Erase(1, 3)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(1, 3)\n", v.data, v.finish, n)
	n = v.Erase(1, 4)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(1, 4)\n", v.data, v.finish, n)
	n = v.Erase(4, 1)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(4, 1)\n", v.data, v.finish, n)
	n = v.Erase(1, 0)
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Erase(1, 0)\n", v.data, v.finish, n)
}

// go test -v vector.go vector_test.go -test.run TestClear
func TestClear(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	v.PushBack("b")
	v.PushBack("c")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	v.Clear()
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
}

// go test -v vector.go vector_test.go -test.run TestInsert
func TestInsert(t *testing.T) {
	v := NewVector()
	v.PushBack("a")
	v.PushBack("b")
	v.PushBack("c")
	v.PushBack("d")
	v.PushBack("e")
	v.PushBack("f")
	fmt.Printf("v.data:%v, v.finish:%v\n", v.data, v.finish)
	n := v.Insert(0, "0")
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Insert(0, \"0\")\n", v.data, v.finish, n)
	n = v.Insert(2, "2", "3", "4")
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Insert(2, \"2\", \"3\", \"4\")\n", v.data, v.finish, n)
	n = v.Insert(9, "9")
	fmt.Printf("v.data:%v, v.finish:%v, %v = v.Insert(9, \"9\")\n", v.data, v.finish, n)
}

// go test -v vector.go vector_test.go -test.run TestCopy
func TestCopy(t *testing.T) {
	v1 := NewVector()
	v1.PushBack("a")
	v1.PushBack("b")
	v1.PushBack("c")

	v2 := NewVector()
	v2.PushBack(1)
	v2.PushBack(2)
	v2.PushBack(3)

	fmt.Printf("before:\n")
	fmt.Printf("v1.data:%v, v1.finish:%v\n", v1.data, v1.finish)
	fmt.Printf("v2.data:%v, v2.finish:%v\n", v2.data, v2.finish)

	n := v1.Copy(v2, 1, 2, 1)
	fmt.Printf("v1.data:%v, v1.finish:%v, %d = v1.Copy(v2, 1, 2, 1)\n", v1.data, v1.finish, n)
	fmt.Printf("v2.data:%v, v2.finish:%v, %d = v1.Copy(v2, 1, 2, 1)\n", v2.data, v2.finish, n)

	n = v1.Copy(v2, 2, 1, 3)
	fmt.Printf("after v1.Copy(v2, 2, 1, 3):\n")
	fmt.Printf("v1.data:%v, v1.finish:%v, %d = v1.Copy(v2, 2, 1, 3)\n", v1.data, v1.finish, n)
	fmt.Printf("v2.data:%v, v2.finish:%v, %d = v1.Copy(v2, 2, 1, 3)\n", v2.data, v2.finish, n)

	n = v1.Copy(v2, 1, 2, 3)
	fmt.Printf("after v1.Copy(v2, 1, 2, 3):\n")
	fmt.Printf("v1.data:%v, v1.finish:%v, %d = v1.Copy(v2, 1, 2, 3)\n", v1.data, v1.finish, n)
	fmt.Printf("v2.data:%v, v2.finish:%v, %d = v1.Copy(v2, 1, 2, 3)\n", v2.data, v2.finish, n)
}

// go test -v vector.go vector_test.go -test.run TestSwap
func TestSwap(t *testing.T) {
	v1 := NewVector()
	v1.PushBack("a1")
	v1.PushBack("b1")
	v1.PushBack("c1")

	v2 := NewVector()
	v2.PushBack("a2")
	v2.PushBack("b2")
	v2.PushBack("c2")
	v2.PushBack("d2")

	fmt.Printf("before:\n")
	fmt.Printf("v1.data:%v, v1.finish:%v\n", v1.data, v1.finish)
	fmt.Printf("v2.data:%v, v2.finish:%v\n", v2.data, v2.finish)

	v1.Swap(v2, 1, 3, 0)
	fmt.Printf("after v1.Swap(v2):\n")
	fmt.Printf("v1.data:%v, v1.finish:%v\n", v1.data, v1.finish)
	fmt.Printf("v2.data:%v, v2.finish:%v\n", v2.data, v2.finish)
}
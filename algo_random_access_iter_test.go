package gocore

import (
	"fmt"
	"testing"
)

// 随机访问迭代器
type RIter struct {
	data interface{}

	prev RandomAccessIter
	next RandomAccessIter
}

func (ri1 *RIter) Equal(ri2 RandomAccessIter) bool {
	return ri1 == ri2
}

func (ri1 *RIter) Get() interface{} {
	return ri1.data
}

func (ri1 *RIter) Set(data interface{}) {
	ri1.data = data
}

func (ri1 *RIter) Next() RandomAccessIter {
	return ri1.next
}

func (ri1 *RIter) Prev() RandomAccessIter {
	return ri1.prev
}

func (ri1 *RIter) Add(d Distance) RandomAccessIter {

	return nil
}

func (ri1 *RIter) Sub(ri2 RandomAccessIter) Distance {
	var back Distance = 0
	for ; ri2 != nil && !ri2.Equal(ri1); ri2 = ri2.Next() {
		back += 1
	}
	return back
}

// go test -v iterator.go algo_random_access_iter.go algo_random_access_iter_test.go -test.run TestSwapRangeRandomAccessIter
func TestSwapRangeRandomAccessIter(t *testing.T) {
	ri1 := &RIter{data: 1}
	ri2 := &RIter{data: 2}
	ri3 := &RIter{data: 3}
	ri4 := &RIter{data: 4}

	ri1.prev = ri4
	ri1.next = ri2

	ri2.prev = ri1
	ri2.next = ri3

	ri3.prev = ri2
	ri3.next = ri4

	ri4.prev = ri3
	ri4.next = ri1

	fmt.Printf("before:\n")
	fmt.Printf("ri1:%v, ri2:%v, ri3:%v, ri4:%v\n", ri1.Get(), ri2.Get(), ri3.Get(), ri4.Get())
	ri := SwapRangeRandomAccessIter(ri1, ri3, ri3)
	fmt.Printf("after SwapRangeRandomAccessIter(ri1, ri3, ri3):\n")
	fmt.Printf("ri1:%v, ri2:%v, ri3:%v, ri4:%v, ri:%v\n", ri1.Get(), ri2.Get(), ri3.Get(), ri4.Get(), ri.Get())
}

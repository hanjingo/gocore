package gocore

import (
	"fmt"
	"testing"
)

// 双向迭代器
type BIter struct {
	data interface{}

	prev BidirectionalIterator
	next BidirectionalIterator
}

func (bi1 *BIter) Equal(bi2 BidirectionalIterator) bool {
	return bi1 == bi2
}

func (bi *BIter) Get() interface{} {
	return bi.data
}

func (bi *BIter) Set(data interface{}) {
	bi.data = data
}

func (bi *BIter) Next() BidirectionalIterator {
	return bi.next
}

func (bi *BIter) Prev() BidirectionalIterator {
	return bi.prev
}

// go test -v iterator.go algo.go algo_test.go -test.run TestSwapRangeBidirectionalIter
func TestSwapRangeBidirectionalIter(t *testing.T) {
	ri1 := &BIter{data: 1}
	ri2 := &BIter{data: 2}
	ri3 := &BIter{data: 3}
	ri4 := &BIter{data: 4}
	
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
	ri := SwapRangeBidirectionalIter(ri1, ri3, ri2)
	fmt.Printf("after SwapRangeBidirectionalIter(ri1, ri3, ri2):\n")
	fmt.Printf("ri1:%v, ri2:%v, ri3:%v, ri4:%v, ri:%v\n", ri1.Get(), ri2.Get(), ri3.Get(), ri4.Get(), ri.Get())
}
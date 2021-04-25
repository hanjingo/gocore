package gocore

/**
 * Vector
 *
 * Size() uint64
 * Empty() bool
 * At(idx VectorIter) interface{}
 * Front() interface{}
 * Back() interface{}
 * Range(first, last VectorIter) []interface{}
 * PopBack()
 * PopFront()
 * PushBack(arg interface{})
 * PushFront(arg interface{})
 * Erase(first, last VectorIter) uint64
 * Clear()
 **/

type VectorIter struct {
	v    *Vector
	prev *VectorIter
	next *VectorIter
}

func (vi1 *VectorIter) Equal(RandomAccessIter) bool {
	return true
}

func (vi *VectorIter) Get() interface{} {
	return nil
}

func (vi *VectorIter) Set(data interface{}) {

}

func (vi *VectorIter) Next() RandomAccessIter {
	return vi.next
}

func (vi *VectorIter) Prev() RandomAccessIter {
	return vi.prev
}

func (vi1 *VectorIter) Add(d Distance) RandomAccessIter {
	return nil
}

func (vi1 *VectorIter) Sub(RandomAccessIter) Distance {
	return 0
}

type Vector struct {
	data []interface{}
	first *VectorIter
	last *VectorIter
	len uint64
}

func (v *Vector) Init() {
	v.data = []interface{}{}
	v.first.next = v.last
	v.first.prev = v.last
	v.len = 0
}

func (v *Vector) Clear() {
	v.Init()
}
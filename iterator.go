package gocore

// 距离
type Distance uint64

// 输入迭代器
type InputIterator interface {
	Equal(InputIterator) bool
	Get() interface{}
	Next() InputIterator
}

// 输出迭代器
type OutputIterator interface {
	Set(data interface{})
	Next() OutputIterator
}

// 向前迭代器
type ForwardIterator interface {
	Equal(ForwardIterator) bool
	Get() interface{}
	Set(data interface{})
	Next() ForwardIterator
}

// 双向迭代器
type BidirectionalIterator interface {
	Equal(BidirectionalIterator) bool
	Get() interface{}
	Set(data interface{})
	Next() BidirectionalIterator
	Prev() BidirectionalIterator
}

// 随机访问迭代器
type RandomAccessIter interface {
	Equal(RandomAccessIter) bool
	Get() interface{}
	Set(data interface{})
	Next() RandomAccessIter
	Prev() RandomAccessIter
	Add(d Distance) RandomAccessIter
	Sub(RandomAccessIter) Distance
}
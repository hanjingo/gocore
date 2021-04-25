package gocore

// 交换双向迭代器
func SwapBidirectionalIter(a, b BidirectionalIterator) {
	tmp := a.Get()
	a.Set(b.Get())
	b.Set(tmp)
}

// 范围交换双向迭代器; 将容器1的[first1, last1)与容器2的[first2, last1-first1)的部分
func SwapRangeBidirectionalIter(first1, last1, first2 BidirectionalIterator) BidirectionalIterator {
	for ; !first1.Equal(last1); {
		SwapBidirectionalIter(first1, first2)
		first1 = first1.Next()
		first2 = first2.Next()
	}
	return first2
}
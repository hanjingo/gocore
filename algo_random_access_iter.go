package gocore

// 交换随机访问迭代器
func SwapRandomAccessIter(a, b RandomAccessIter) {
	tmp := a.Get()
	a.Set(b.Get())
	b.Set(tmp)
}

// 范围交换随机访问迭代器; 将容器1的[first1, last1)与容器2的[first2, last1-first1)的部分
func SwapRangeRandomAccessIter(first1, last1, first2 RandomAccessIter) RandomAccessIter {
	for ; !first1.Equal(last1); {
		SwapRandomAccessIter(first1, first2)
		first1 = first1.Next()
		first2 = first2.Next()
	}
	return first2
}

// 复制随机访问迭代器; 将容器1的[first, last)复制到容器2的[pos, last-first)
func CopyRandomAccessIter(first, last, pos RandomAccessIter) RandomAccessIter {
	for n := last.Sub(first); n > 0 && pos != nil && first != nil; n-- {
		pos.Set(first.Get())
		first = first.Next()
		pos = pos.Next()
	}
	return pos
}

// 复制非重复的元素; 
func UniqueCopyRandomAccessIter(first, last, pos RandomAccessIter) RandomAccessIter {
	if first.Equal(last) {
		return pos
	}
	value := first.Get()
	pos.Set(value)
	for !first.Equal(last) {
		if value != first.Get() {
			value = first.Get()
			pos = pos.Next()
			pos.Set(value)
		}
		first = first.Next()
	}
	return pos.Next()
}

// 移除[first, last)的重复元素
func UniqueRandomAccessIter(first, last RandomAccessIter) RandomAccessIter {
	return UniqueCopyRandomAccessIter(first, last, first)
}

// 翻转区间[first, last)
func ReverseRandomAccessIter(first, last RandomAccessIter) {
	for {
		last = last.Prev()
		if first.Equal(last) || first.Equal(last){
			return
		} else {
			SwapRandomAccessIter(first, last)
			first = first.Next()
		}
	}
}

// 将容器的区段[first, middle)和区段[middle, last)的元素互换
func RotateRandomAccessIter(first, middle, last RandomAccessIter) RandomAccessIter {
	if first.Equal(middle) {
		return last
	}
	if last.Equal(middle) {
		return first
	}

	ReverseRandomAccessIter(first, middle)
	ReverseRandomAccessIter(middle, last)

	for !first.Equal(middle) && !middle.Equal(last) {
		last = last.Prev()
		SwapRandomAccessIter(first, last)
		first = first.Next()
	}

	if first.Equal(middle) {
		ReverseRandomAccessIter(middle, last)
		return last
	} else {
		ReverseRandomAccessIter(first, middle)
		return first
	}
}

// 查找容器中区段[first, last)的值value
func FindRandomAccessIter(first, last RandomAccessIter, value interface{}) RandomAccessIter {
	for !first.Equal(last) && first.Get() != value {
		first = first.Next()
	}
	return first
}

// 删除容器的区段[first, last)的与value相等的元素
func RemoveCopyRandomAccessIter(first, last, pos RandomAccessIter, value interface{}) RandomAccessIter {
	for ; !first.Equal(last); first = first.Next() {
		if first.Get() != value {
			pos.Set(first.Get())
			pos = pos.Next()
		}
	}
	return pos
}

// 删除容器的区段[first, last)的值value
func RemoveRandomAccessIter(first, last RandomAccessIter, value interface{}) RandomAccessIter {
	first = FindRandomAccessIter(first, last, value)
	if first.Equal(last) {
		return first
	} else {
		back := RemoveCopyRandomAccessIter(first.Next(), last, first, value)
		first = first.Next()
		return back
	}
}
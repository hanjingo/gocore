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
 * Insert(position VectorIter, datas ...interface{}) uint64
 * Copy(v2 *Vector, first2, last2, position1 VectorIter) uint64
 * Swap(v2 *Vector, first2, last2, position1 VectorIter) uint64
 **/

type VectorIter uint64
type Vector struct {
	data   []interface{}
	finish VectorIter
}

func NewVector() *Vector {
	return &Vector{
		data:   []interface{}{},
		finish: 0,
	}
}

func (v *Vector) Size() uint64 {
	return uint64(v.finish)
}

func (v *Vector) Empty() bool {
	return v.finish == 0
}

func (v *Vector) At(idx VectorIter) interface{} {
	if idx >= v.finish {
		return nil
	}
	return v.data[idx]
}

func (v *Vector) Front() interface{} {
	if v.Empty() {
		return nil
	}
	return v.data[0]
}

func (v *Vector) Back() interface{} {
	if v.Empty() {
		return nil
	}
	return v.data[v.finish-1]
}

// Range (first, last]
func (v *Vector) Range(first, last VectorIter) []interface{} {
	if v.Empty() || first > last || last >= v.finish{
		return nil
	}
	back := make([]interface{}, last - first)
	copy(back, v.data[first:last])
	return back
}

func (v *Vector) PopBack() {
	if v.Empty() {
		return
	}
	v.data = v.data[:v.finish-1]
	v.finish -= 1
}

func (v *Vector) PopFront() {
	if v.Empty() {
		return
	}
	v.data = v.data[1:]
	v.finish -= 1
}

func (v *Vector) PushBack(arg interface{}) {
	v.data = append(v.data, arg)
	v.finish += 1
}

func (v *Vector) PushFront(arg interface{}) {
	v.data = append([]interface{}{arg}, v.data...)
	v.finish += 1
}

// Erase (first, last], return num of erased elements
func (v *Vector) Erase(first, last VectorIter) uint64 {
	if last < first || last > v.finish {
		return 0
	}
	v.data = append(v.data[:first], v.data[last:]...)
	v.finish = VectorIter(len(v.data))
	return uint64(last - first)
}

func (v *Vector) Clear() {
	v.data = []interface{}{}
	v.finish = 0
}

func (v *Vector) Insert(position VectorIter, datas ...interface{}) uint64 {
	if position >= v.finish {
		return 0
	}
	size := VectorIter(len(datas))
	v.data = append(v.data[:position], append(datas, v.data[position:]...)...)
	v.finish += size
	return uint64(size)
}

// copy data=v2(first2, last2] to v1 -> ...position1(data]...
func (v1 *Vector) Copy(v2 *Vector, first2, last2, position1 VectorIter) uint64 {
	if datas := v2.Range(first2, last2); datas != nil {
		size := v1.Insert(position1, datas...)
		v1.finish += VectorIter(size)
		return size
	}
	return 0
}

// swap data=v2(first2, last2] to v1 -> ...position1(data]...
func (v1 *Vector) Swap(v2 *Vector, first2, last2, position1 VectorIter) uint64 {
	if datas2 := v2.Range(first2, last2); datas2 != nil {
		datas2Last := VectorIter(len(datas2))
		if datas1 := v1.Range(position1, datas2Last); datas1 != nil {
			// erase and insert vector2
			v2.Erase(first2, last2)
			v2.Insert(first2, datas1...)

			// erase and insert vector1
			data1Last := VectorIter(len(datas1))
			v1.Erase(position1, data1Last)
			v2.Insert(position1, datas2...)

			// set finish
			v1.finish = VectorIter(len(v1.data))
			v2.finish = VectorIter(len(v2.data))
		}
	}
	return 0
}

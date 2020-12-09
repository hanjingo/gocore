package core

import "container/list"

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	return &Queue{
		data: list.New(),
	}
}

func (q *Queue) Front() interface{} {
	if e := q.data.Front(); e != nil {
		return e.Value
	}
	return nil
}

func (q *Queue) Back() interface{} {
	if e := q.data.Back(); e != nil {
		return e.Value
	}
	return nil
}

func (q *Queue) Push(data interface{}) {
	q.data.PushBack(data)
}

func (q *Queue) Pop() interface{} {
	if e := q.data.Front(); e != nil {
		q.data.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) Size() int {
	return q.data.Len()
}

func (q *Queue) Empty() bool {
	return q.data.Len() == 0
}

func (q *Queue) Swap(other *Queue) {
	tmp := q.data
	q.data = other.data
	other.data = tmp
}

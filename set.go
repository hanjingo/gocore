package gocore

var void struct{}

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(map[interface{}]struct{})
}

func (s Set) Has(e interface{}) bool {
	_, ok := s[e]
	return ok
}

func (s Set) Add(e interface{}) {
	s[e] = void
}

func (s Set) Del(e interface{}) {
	delete(s, e)
}

func (s Set) Range(f func(k interface{})) {
	for e, _ := range s {
		f(e)
	}
}

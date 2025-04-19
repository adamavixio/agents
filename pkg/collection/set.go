package collection

type Set[T comparable] interface {
	Put(T)
	Has(T) bool
	Del(T)
}

type set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(set[T])
}

func (s set[T]) Put(v T)      { s[v] = struct{}{} }
func (s set[T]) Has(v T) bool { _, ok := s[v]; return ok }
func (s set[T]) Del(v T)      { delete(s, v) }

package collection

type Map[K comparable, V any] interface {
	Put(K, V) bool
	Has(K) bool
	Get(K) (V, bool)
	Del(K) bool
}

type mapped[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return make(mapped[K, V])
}

func (m mapped[K, V]) Put(k K, v V) bool {
	if m.Has(k) {
		return false
	}
	m[k] = v
	return true
}

func (m mapped[K, V]) Has(k K) bool {
	_, ok := m[k]
	return ok
}

func (m mapped[K, V]) Get(k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}

func (m mapped[K, V]) Del(k K) bool {
	if m.Has(k) {
		return false
	}
	delete(m, k)
	return true
}

package gmap

import "sync"

type SafeMap[K comparable, V any] struct {
	m sync.Map
}

func (s *SafeMap[K, V]) Load(k K) (v V, ok bool) {
	if v, ok := s.m.Load(k); ok {
		return v.(V), true
	}
	return v, false
}

func (s *SafeMap[K, V]) LoadAndDelete(k K) (v V, loaded bool) {
	if v, ok := s.m.LoadAndDelete(k); ok {
		return v.(V), true
	}
	return v, false
}

func (s *SafeMap[K, V]) LoadOrStore(k K, v V) (r V, loaded bool) {
	if v, ok := s.m.LoadOrStore(k, v); ok {
		return v.(V), true
	}
	return r, false
}

func (s *SafeMap[K, V]) CompareAndDelete(k K, old V) (deleted bool) {
	return s.m.CompareAndDelete(k, old)
}

func (s *SafeMap[K, V]) CompareAndSwap(k K, old V, new V) (swapped bool) {
	return s.m.CompareAndSwap(k, old, new)
}

func (s *SafeMap[K, V]) Delete(k K) {
	s.m.Delete(k)
}

func (s *SafeMap[K, V]) Clear() {
	s.m.Clear()
}

func (s *SafeMap[K, V]) Store(k K, v V) {
	s.m.Store(k, v)
}

func (s *SafeMap[K, V]) Swap(k K, v V) (old V, loaded bool) {
	if old, ok := s.m.Swap(k, v); ok {
		return old.(V), true
	}
	return old, false
}

package sync

import (
	"iter"
	"sync"
)

type SyncMap[K, V any] struct {
	m sync.Map
}

type ISyncMap[K, V any] interface {
	Store(key K, val V)
	Load(key K) (V, bool)
	LoadOrStore(key K, val V) (V, bool)
	Delete(key K)
	LoadAndDelete(key K) (V, bool)
	Iter() iter.Seq2[K, V]
}

func (s *SyncMap[K, V]) Store(key K, val V) {
	s.m.Store(key, val)
}

func (s *SyncMap[K, V]) retVal(val any, ok bool) (V, bool) {
	if ok {
		switch val.(type) {
		case nil:
			return *new(V), ok
		default:
			return val.(V), ok
		}
	} else {
		return *new(V), ok
	}
}

func (s *SyncMap[K, V]) Load(key K) (V, bool) {
	val, ok := s.m.Load(key)
	return s.retVal(val, ok)

}

func (s *SyncMap[K, V]) LoadOrStore(key K, val V) (V, bool) {
	actual, loaded := s.m.LoadOrStore(key, val)
	return s.retVal(actual, loaded)
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	val, loaded := s.m.LoadAndDelete(key)
	return s.retVal(val, loaded)
}

type Entry[K, V any] struct {
	key K
	val V
}

func (s *SyncMap[K, V]) Iter() iter.Seq2[K, V] {
	var entries []Entry[K, V]
	s.m.Range(func(key, value any) bool {
		entries = append(entries, Entry[K, V]{key.(K), value.(V)})
		return true
	})

	return func(yield func(K, V) bool) {
		for _, entry := range entries {
			if !yield(entry.key, entry.val) {
				return
			}
		}
	}
}
